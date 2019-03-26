package host

import (
	"fmt"
	"net"
)

type Host struct {
	name   string `json:"name"`
	ipAddr net.IP `json:"ip_addr"`
}

func New(name, ipAddr string) (*Host, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return nil, fmt.Errorf("IP Address %v is invalid for host %v", ipAddr, name)
	}
	return &Host{name, ip}, nil
}

func (h *Host) Name() string {
	return h.name
}

func (h *Host) IPAddr() net.IP {
	return h.ipAddr
}
