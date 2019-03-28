package network

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBridgeIPNetHasIP(t *testing.T) {
	b, _ := NewBridge("fake-br")
	i := NewIPLink()
	ip, ipnet, _ := net.ParseCIDR("1.2.3.4/24")

	i.setIfaceIP(b.Name(), ip, ipnet.Mask)

	brIP, brIPMask, err := GetBridgeIPNet(b.Name())
	assert.Equal(t, ip, brIP)
	assert.Equal(t, ipnet.Mask, brIPMask)
	assert.NoError(t, err)
}

func TestGetBridgeIPNetBridgeDoesnotExist(t *testing.T) {
	brIP, brIPMask, err := GetBridgeIPNet("fake-bar-br")
	assert.Equal(t, net.IP{}, brIP)
	assert.Equal(t, net.IPMask{}, brIPMask)
	assert.Error(t, err)
}

func TestGetBridgeIPNetHasNoIP(t *testing.T) {
	b, _ := NewBridge("fake-foo-br")

	brIP, brIPMask, err := GetBridgeIPNet(b.Name())
	assert.Equal(t, net.IP{}, brIP)
	assert.Equal(t, net.IPMask{}, brIPMask)
	assert.EqualError(t, err, fmt.Sprintf("Host bridge %q doesn't have IP", b.Name()))
}