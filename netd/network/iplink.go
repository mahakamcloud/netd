package network

import (
	"fmt"
	"net"
	"strconv"

	"github.com/mahakamcloud/netd/cmdrunner"
	"github.com/mahakamcloud/netd/logger"
)

type ipLink struct {
	runner cmdrunner.CmdRunner
}

func NewIPLink() *ipLink {
	return &ipLink{
		runner: cmdrunner.New(),
	}
}

func (i *ipLink) createBridge(name string) error {
	if i.ifaceExists(name) {
		return nil
	}
	output, err := i.runner.CombinedOutput("ip", "link", "add", "name", name, "type", "bridge")
	if err != nil {
		logger.Log.Errorf("Error creating bridge %v: %v", name, err)
		return err
	}
	logger.Log.Debugf("Bridge %v created: %v", name, output)
	return nil
}

func (i *ipLink) createTapDev(name string) error {
	if i.ifaceExists(name) {
		return nil
	}
	output, err := i.runner.CombinedOutput("ip", "tuntap", "add", "dev", name, "mode", "tap")
	if err != nil {
		logger.Log.Errorf("Error creating tap device %v: %v", name, err)
		return err
	}
	logger.Log.Debugf("Tap device %v created: %v", name, output)
	return nil
}

func (i *ipLink) createGRETapDev(name string, localIP, remoteIP net.IP, key int) error {
	if i.ifaceExists(name) {
		return nil
	}
	output, err := i.runner.CombinedOutput("ip", "link", "add", name, "type", "gretap", "remote", remoteIP.String(), "local", localIP.String(), "key", strconv.Itoa(key))
	if err != nil {
		logger.Log.Errorf("Error creating GRE tap device %v: %v", name, err)
		return err
	}
	logger.Log.Debugf("GRE tap device %v created: %v", name, output)
	return nil
}

func (i *ipLink) addIfaceToBridge(ifaceName, brName string) error {
	output, err := i.runner.CombinedOutput("ip", "link", "set", "dev", ifaceName, "master", brName)
	if err != nil {
		logger.Log.Errorf("Error adding interface %s to bridge %s: %v", ifaceName, brName, err)
		return err
	}
	logger.Log.Debugf("Interface %s added to bridge %s: %v", ifaceName, brName, output)
	return nil
}

func (i *ipLink) ifaceExists(name string) bool {
	output, err := i.runner.CombinedOutput("ip", "link", "show", "dev", name)
	if err != nil {
		logger.Log.Debugf("Interface %v does not exists: %v", name, err)
		return false
	}
	logger.Log.Debugf("Interface %v exists: %v", name, output)
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
