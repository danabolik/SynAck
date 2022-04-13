package workers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type DialerStub struct {
	mock.Mock
}

type DeliveryStub struct {
}

type ProducerStub struct {
	psChan chan int
}

func (ps ProducerStub) WritePsToChan(psChan chan int) {
	for i := 1; i <= cap(psChan); i++ {
		psChan <- i
	}
	close(psChan)
}

func (ps ProducerStub) GetCountPorts() int {
	return 5
}

func (ds DeliveryStub) GetTcpNetwork() string {
	return "tcp"
}

func (d *DialerStub) DialPort(network, addr string, p int) int {
	args := d.Called(network, addr, p)
	return args.Int(0)
}

func TestScanPorts(t *testing.T) {
	addr := "scanme.nmap.org"
	grt := 5

	delivery := DeliveryStub{}

	dialer := &DialerStub{}
	dialer.On("DialPort", delivery.GetTcpNetwork(), addr, mock.Anything).Once().Return(1)
	dialer.On("DialPort", delivery.GetTcpNetwork(), addr, mock.Anything).Once().Return(2)
	dialer.On("DialPort", delivery.GetTcpNetwork(), addr, mock.Anything).Once().Return(3)
	dialer.On("DialPort", delivery.GetTcpNetwork(), addr, mock.Anything).Once().Return(4)
	dialer.On("DialPort", delivery.GetTcpNetwork(), addr, mock.Anything).Once().Return(5)

	producer := ProducerStub{}
	psChan := make(chan int, producer.GetCountPorts())

	w := Worker{Decorator: dialer, Delivery: delivery, Producer: &producer}

	producer.WritePsToChan(psChan)
	result := w.ScanPorts(addr, grt)

	exp := []int{1, 2, 3, 4, 5}
	for _, act := range result {
		assert.Contains(t, exp, act)
	}
}

func TestScanPortsWhenEmpty(t *testing.T) {
	addr := "scanme.nmap.org"
	grt := 5

	producer := ProducerStub{}
	cntPs := producer.GetCountPorts()

	delivery := DeliveryStub{}

	dialer := &DialerStub{}
	dialer.On("DialPort", delivery.GetTcpNetwork(), addr, mock.Anything).Times(cntPs).Return(0)
	psChan := make(chan int, cntPs)

	w := Worker{Decorator: dialer, Delivery: delivery, Producer: &producer}

	producer.WritePsToChan(psChan)
	result := w.ScanPorts(addr, grt)

	assert.Empty(t, result)
}
