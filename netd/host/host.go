package host

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/mahakamcloud/netd/client"
	"github.com/mahakamcloud/netd/config"
)

const (
	MahakamHostRegistrationAPI = "/bare-metal-hosts"
	MahakamHostBaseUrl         = "http://%s:%d/v1"
)

type Host struct {
	Name       string `json:"name"`
	IPAddr     net.IP `json:"ip"`
	IPMaskSize int    `json:"ipmask_size"`
}

func New(name string, ip net.IP, ipMask net.IPMask) *Host {
	mask, _ := ipMask.Size()
	return &Host{name, ip, mask}
}

func (h *Host) Register(c *client.Client) error {
	j, err := json.Marshal(h)
	if err != nil {
		return err
	}

	url := fmt.Sprintf(MahakamHostBaseUrl+MahakamHostRegistrationAPI, config.MahakamIP(), config.MahakamPort())
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	_, err = c.DoRequest(req)
	if err != nil {
		return err
	}
	return nil
}
