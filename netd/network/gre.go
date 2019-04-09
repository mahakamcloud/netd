package network

import (
	"net"

	log "github.com/sirupsen/logrus"
)

type GRE struct {
	name     string
	localIP  net.IP
	remoteIP net.IP
	key      int
	iplink   *ipLink
}

func NewGRE(name string, localIP, remoteIP net.IP, key int) *GRE {
	return &GRE{
		name:     name,
		localIP:  localIP,
		remoteIP: remoteIP,
		key:      key,
		iplink:   NewIPLink(),
	}
}

func (g *GRE) Create(bridgeName string) error {
	err := g.iplink.createGRETapDev(g.name, g.localIP, g.remoteIP, g.key)
	if err != nil {
		log.Errorf("Error creating GRE tap device %v with key %d: %v", g.name, g.key, err)
		return err
	}

	err = g.iplink.addIfaceToBridge(g.name, bridgeName)
	if err != nil {
		log.Errorf("Error adding port %v to bridge %v: %v", g.name, bridgeName, err)
		return err
	}

	_, err = g.iplink.setIfaceUp(g.name)
	if err != nil {
		log.Errorf("Error bringing tap interface %q up: %v", g.name, err)
		return err
	}

	return nil
}

func (g *GRE) Name() string {
	return g.name
}
