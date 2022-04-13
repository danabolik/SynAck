package delivery

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttp_GetNetwork(t *testing.T) {
	assert.Equal(t, "tcp", Http{}.GetTcpNetwork())
}
