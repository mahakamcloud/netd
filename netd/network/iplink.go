package network

import (
	"github.com/mahakamcloud/mahakam/pkg/cmd_runner"
	log "github.com/sirupsen/logrus"
)

type IPUtil struct {
	runner cmd_runner.CmdRunner
}

func NewIPUtil() *IPUtil {
	return &IPUtil{
		runner: cmd_runner.New(),
	}
}

func (ipu *IPUtil) CreateTapDev(name string) (string, error) {
	if ipu.tapDevExists(name) {
		return "", nil
	}
	return ipu.runner.CombinedOutput("ip", "tuntap", "add", "dev", name, "mode", "tap")
}

func (ipu *IPUtil) tapDevExists(name string) bool {
	output, err := ipu.runner.CombinedOutput("ip", "link", "show", "dev", name)
	if err != nil {
		log.Debugf("Tap device %v does not exists: %v", name, err)
		return false
	}
	log.Debugf("Tap device %v exists: %v", name, output)
	return true
}

func (ipu *IPUtil) SetIfaceUp(name string) (string, error) {
	args := []string{"link", "set", "dev", name, "up"}
	return ipu.runner.CombinedOutput("ip", args...)
}
