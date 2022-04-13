package decorators

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"time"
)

type ConnStub struct {
	isClosed bool
}

func (c ConnStub) Read(b []byte) (n int, err error) {
	return 0, nil
}
func (c ConnStub) Write(b []byte) (n int, err error) {
	return 0, nil
}
func (c ConnStub) Close() error {
	if c.isClosed {
		return errors.New("failed close connection")
	}
	return nil
}
func (c ConnStub) LocalAddr() net.Addr {
	return nil
}
func (c ConnStub) RemoteAddr() net.Addr {
	return nil
}
func (c ConnStub) SetDeadline(t time.Time) error {
	return nil
}
func (c ConnStub) SetReadDeadline(t time.Time) error {
	return nil
}
func (c ConnStub) SetWriteDeadline(t time.Time) error {
	return nil
}

type TestDialer struct {
	isOpen bool
	conn   ConnStub
}

func (d TestDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	if !d.isOpen {
		return d.conn, errors.New("failed connection")
	}
	return d.conn, nil
}

type dialAllDataProvider struct {
	dialer   TestDialer
	network  string
	address  string
	port     int
	openPort int
}

func TestNetDial_DialPort(t *testing.T) {
	provider := []dialAllDataProvider{
		{
			TestDialer{isOpen: true, conn: ConnStub{isClosed: false}},
			"tcp",
			"scanme.nmap.org",
			80,
			80,
		},
		{
			TestDialer{isOpen: false, conn: ConnStub{isClosed: false}},
			"tcp",
			"scanme.nmap.org",
			443,
			0,
		},
	}
	for _, p := range provider {
		ps := NetDecorator{p.dialer}.DialPort(p.network, p.address, p.port)
		assert.Equal(t, p.openPort, ps)
	}
}

type PanicDataProvider struct {
	dialer  TestDialer
	network string
	address string
	port    int
}

func TestNetDial_DialPortWhenPanic(t *testing.T) {
	p := PanicDataProvider{
		TestDialer{isOpen: true, conn: ConnStub{isClosed: true}},
		"tcp",
		"scanme.nmap.org",
		664,
	}
	assert.Panics(t, func() {
		NetDecorator{p.dialer}.DialPort(p.network, p.address, p.port)
	})
}
