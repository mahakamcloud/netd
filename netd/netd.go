package netd

import (
	"fmt"
	"net"
	"os"

	"github.com/mahakamcloud/netd/client"
	"github.com/mahakamcloud/netd/config"
	"github.com/mahakamcloud/netd/netd/host"
	log "github.com/sirupsen/logrus"
)

func Register() {
	hostName, err := os.Hostname()
	if err != nil {
		log.Errorf("error getting hostname: %v", err)
		return
	}

	ip, ipMask, err := getBridgeIPNet(config.HostBridgeName())
	if err != nil {
		log.Errorf("error getting ip network: %v", err)
		return
	}

	h, err := host.New(hostName, ip, ipMask)
	if err != nil {
		log.Errorf("error initializing host: %v", err)
		return
	}

	err = h.Register(&client.Client{})
	if err != nil {
		log.Errorf("error registering host to mahakam: %v", err)
		return
	}
}

func getBridgeIPNet(bridgeName string) (net.IP, net.IPMask, error) {
	iface, err := net.InterfaceByName(bridgeName)
	if err != nil {
		return net.IP{}, net.IPMask{}, err
	}

	if addrs, _ := iface.Addrs(); len(addrs) > 0 {
		ip, ipNet, _ := net.ParseCIDR(addrs[0].String())
		return ip, ipNet.Mask, nil
	}
	return nil, nil, fmt.Errorf("host bridge %q doesn't have IP", bridgeName)
}
