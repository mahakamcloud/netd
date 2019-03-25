package network

import (
	"github.com/mahakamcloud/mahakam/pkg/cmd_runner"
	log "github.com/sirupsen/logrus"
)

type IPLink struct {
	runner cmd_runner.CmdRunner
}

func NewIPLink() *IPLink {
	return &IPLink{
		runner: cmd_runner.New(),
	}
}

func (i *IPLink) CreateTapDev(name string) (string, error) {
	if i.tapDevExists(name) {
		return "", nil
	}
	return i.runner.CombinedOutput("ip", "tuntap", "add", "dev", name, "mode", "tap")
}

func (i *IPLink) tapDevExists(name string) bool {
	output, err := i.runner.CombinedOutput("ip", "link", "show", "dev", name)
	if err != nil {
		log.Debugf("Tap device %v does not exists: %v", name, err)
		return false
	}
	log.Debugf("Tap device %v exists: %v", name, output)
	return true
}

func (i *IPLink) SetIfaceUp(name string) (string, error) {
	args := []string{"link", "set", "dev", name, "up"}
	return i.runner.CombinedOutput("ip", args...)
}
