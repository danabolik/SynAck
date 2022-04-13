package producers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerator_WritePsToChan(t *testing.T) {
	cntPs := 10
	psChan := make(chan int, cntPs)

	g := Generator{}
	g.WritePsToChan(psChan)
	for i := 1; i <= cntPs; i++ {
		assert.Equal(t, i, <-psChan)
	}
}

func TestGenerator_GetCountPorts(t *testing.T) {
	assert.Equal(t, 65536, Generator{}.GetCountPorts())
}
