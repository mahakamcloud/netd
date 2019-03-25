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

func (i *IPLink) createTapDev(name string) error {
	if i.tapDevExists(name) {
		return "", nil
	}
	output, err := i.runner.CombinedOutput("ip", "tuntap", "add", "dev", name, "mode", "tap")
	if err != nil {
		log.Errorf("Error creating tap device %v: %v", name, err)
		return err
	}
	log.Debugf("Tap device %v created: %v", name, output)
	return nil
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

func (i *IPLink) setIfaceUp(name string) (string, error) {
	args := []string{"link", "set", "dev", name, "up"}
	return i.runner.CombinedOutput("ip", args...)
}
