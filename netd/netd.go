package netd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/mahakamcloud/netd/config"
	"github.com/mahakamcloud/netd/mahakamclient"
	"github.com/mahakamcloud/netd/netd/host"
	"github.com/mahakamcloud/netd/netd/network"
)

func Register() error {
	hostName, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("error getting hostname: %v", err)
	}

	ip, ipMask, err := network.GetBridgeIPNet(config.HostBridgeName())
	if err != nil {
		return fmt.Errorf("error getting ip network: %v", err)
	}

	h := host.New(hostName, ip, ipMask)
	err = h.Register(&mahakamclient.Client{})
	if err != nil {
		return fmt.Errorf("error registering host to mahakam: %v", err)
	}

	log.Infof("Successfully registered host %s with IP %s", h.Name, h.IPAddr.String())
	return nil
}
