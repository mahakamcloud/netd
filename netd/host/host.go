package host

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"

	"github.com/mahakamcloud/netd/mahakamclient"
)

type Host struct {
	Name       string `json:"name"`
	IPAddr     net.IP `json:"ip"`
	IPMaskSize int    `json:"ipMask"`
}

func New(name string, ip net.IP, ipMask net.IPMask) *Host {
	mask, _ := ipMask.Size()
	return &Host{name, ip, mask}
}

func (h *Host) Register(c *mahakamclient.Client) error {
	j, err := json.Marshal(h)
	if err != nil {
		return err
	}

	_, err = c.RegisterHost(bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	return nil
}

func (h *Host) String() string {
	return fmt.Sprintf("Name: %s IPAddr: %v IPMaskSize: %d", h.Name, h.IPAddr, h.IPMaskSize)
}
