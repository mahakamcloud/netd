package network

import (
	"fmt"
	"net"

	"github.com/mahakamcloud/netd/cmdrunner"
	log "github.com/sirupsen/logrus"
)

type ipLink struct {
	runner cmdrunner.CmdRunner
}

func NewIPLink() *ipLink {
	return &ipLink{
		runner: cmdrunner.New(),
	}
}

func (i *ipLink) createTapDev(name string) error {
	if i.tapDevExists(name) {
		return nil
	}
	output, err := i.runner.CombinedOutput("ip", "tuntap", "add", "dev", name, "mode", "tap")
	if err != nil {
		log.Errorf("Error creating tap device %v: %v", name, err)
		return err
	}
	log.Debugf("Tap device %v created: %v", name, output)
	return nil
}

func (i *ipLink) tapDevExists(name string) bool {
	output, err := i.runner.CombinedOutput("ip", "link", "show", "dev", name)
	if err != nil {
		log.Debugf("Tap device %v does not exists: %v", name, err)
		return false
	}
	log.Debugf("Tap device %v exists: %v", name, output)
	return true
}

func (i *ipLink) setIfaceUp(name string) (string, error) {
	args := []string{"link", "set", "dev", name, "up"}
	return i.runner.CombinedOutput("ip", args...)
}

func (i *ipLink) setIfaceIP(name string, ip net.IP, ipMask net.IPMask) (string, error) {
	size, _ := ipMask.Size()
	ipCIDR := fmt.Sprintf("%s/%d", ip.String(), size)

	args := []string{"address", "add", ipCIDR, "dev", name}
	return i.runner.CombinedOutput("ip", args...)
}
