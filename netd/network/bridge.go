package network

import (
	"errors"
	"fmt"
	"net"
)

type Bridge struct {
	name string
}

func NewBridge(name string) (*Bridge, error) {
	b := &Bridge{name}
	iplink := NewIPLink()
	err := iplink.createBridge(name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error creating a new bridge %v: %v", name, err))
	}
	return b, nil
}

func (b *Bridge) Name() string {
	return b.name
}

func GetBridgeIPNet(bridgeName string) (net.IP, net.IPMask, error) {
	iface, err := net.InterfaceByName(bridgeName)
	if err != nil {
		return net.IP{}, net.IPMask{}, err
	}

	if addrs, _ := iface.Addrs(); len(addrs) > 0 {
		for _, addr := range addrs {
			if ipAddr, ipnet, _ := net.ParseCIDR(addr.String()); ipAddr.To4() != nil {
				return ipAddr, ipnet.Mask, nil
			}
		}
	}
	return net.IP{}, net.IPMask{}, fmt.Errorf("Host bridge %q doesn't have IP", bridgeName)
}
