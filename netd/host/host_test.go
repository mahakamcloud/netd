package host

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostIsCreated(t *testing.T) {
	h, _ := New("localhost", "127.0.0.1")

	assert.Equal(t, h.Name, "localhost")
	assert.Equal(t, h.IPAddr, net.IPv4(127, 0, 0, 1))
}

func TestHostCreationFailsForInvalidIp(t *testing.T) {
	_, err := New("localhost", "266.266.266.266")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "IP Address 266.266.266.266 is invalid for host localhost")
}
