package host

import "net"

type Host struct {
	name   string `json:"name"`
	ipAddr net.IP `json:"ip_addr"`
}

func (h *Host) Name() string {
	return h.name
}

func (h *Host) IPAddr() net.IP {
	return h.ipAddr
}
