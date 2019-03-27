package netd

import (
	"fmt"
	"net"
	"os"

	"github.com/mahakamcloud/netd/config"
	"github.com/mahakamcloud/netd/mahakamclient"
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

	h := host.New(hostName, ip, ipMask)
	err = h.Register(&mahakamclient.Client{})
	if err != nil {
		log.Errorf("error registering host to mahakam: %v", err)
		return
	}
	log.Infof("Successfully registered host %s with IP %s", h.Name, h.IPAddr.String())
}

func getBridgeIPNet(bridgeName string) (net.IP, net.IPMask, error) {
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
	return nil, nil, fmt.Errorf("host bridge %q doesn't have IP", bridgeName)
}
