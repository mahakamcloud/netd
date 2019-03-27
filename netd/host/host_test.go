package host

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostIsCreated(t *testing.T) {
	ip, ipNet, _ := net.ParseCIDR("1.2.3.4/24")
	h := New("localhost", ip, ipNet.Mask)

	assert.Equal(t, h.name, "localhost")
	assert.Equal(t, h.ipAddr, net.IPv4(1, 2, 3, 4))
	assert.Equal(t, h.ipMask, "24")
}
