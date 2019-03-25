package network

import (
	"net"
	"strconv"

	"github.com/digitalocean/go-openvswitch/ovs"
	log "github.com/sirupsen/logrus"
)

type gre struct {
	name      string
	remoteIP  net.IP
	key       int
	ovsClient *ovs.Client
	iplink    *ipLink
}

func NewGRE(name string, remoteIP net.IP, key int) *gre {
	return &gre{
		name:      name,
		remoteIP:  remoteIP,
		key:       key,
		ovsClient: ovs.New(),
		iplink:    NewIPLink(),
	}
}

func (g *gre) Create(bridgeName string) error {
	err := g.iplink.createTapDev(g.name)
	if err != nil {
		return err
	}

	err = g.ovsClient.VSwitch.AddPort(bridgeName, g.name)
	if err != nil {
		log.Debugf("Error adding port %v to bridge %v: %v", g.name, bridgeName, err)
		return err
	}

	// TODO: OVSwitch set interface is always overwritten. Do not execute if interface set.
	iFaceOptions := ovs.InterfaceOptions{
		Type:     ovs.InterfaceTypeGRE,
		RemoteIP: g.remoteIP.String(),
		Key:      strconv.Itoa(g.key),
	}
	return g.ovsClient.VSwitch.Set.Interface(g.name, iFaceOptions)
}
