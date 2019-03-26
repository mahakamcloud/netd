package host

import (
	"fmt"
	"net"
)

type Host struct {
	Name   string `json:"name"`
	IPAddr net.IP `json:"ip_addr"`
}

func New(name, ipAddr string) (*Host, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return nil, fmt.Errorf("IP Address %v is invalid for host %v", ipAddr, name)
	}
	return &Host{name, ip}, nil
}
