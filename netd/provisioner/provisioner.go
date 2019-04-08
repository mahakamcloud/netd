package provisioner

import (
	"errors"
	"fmt"
	"strings"

	"github.com/libvirt/libvirt-go"
	"github.com/mahakamcloud/netd/netd/cluster"
	"github.com/mahakamcloud/netd/netd/host"
	"github.com/mahakamcloud/netd/netd/network"
)

func generateBridgeName(clusterName string) string {
	var bridgeNamePrefix string
	if len(clusterName) < 12 {
		bridgeNamePrefix = clusterName
	} else {
		bridgeNamePrefix = clusterName[:12]
	}
	return fmt.Sprintf("%s_br", bridgeNamePrefix)
}

func generateGRETunnelName(clusterName, localHostName, remoteHostName string) string {
	// TODO: check for gre tunnel name length
	return fmt.Sprintf("%s_%s_%s", clusterName, localHostName, remoteHostName)
}

func CreateBridge(cl *cluster.Cluster) (*network.Bridge, error) {
	bridgeName := generateBridgeName(cl.Name)
	return network.NewBridge(bridgeName)
}

type GREConnection struct {
	Name       string
	RemoteHost *host.Host
	Status     bool
}

func CreateGREMesh(cl *cluster.Cluster, localhost *host.Host, remotehosts []*host.Host, bridgeName string) ([]*GREConnection, error) {
	greConns := make([]*GREConnection, 0)
	errs := make([]error, 0)

	for _, r := range remotehosts {
		greName := generateGRETunnelName(cl.Name, localhost.Name, r.Name)
		gre := network.NewGRE(greName, localhost.IPAddr, r.IPAddr, cl.GREKey)
		err := gre.Create(bridgeName)

		greConn := &GREConnection{
			Name:       gre.Name(),
			RemoteHost: r,
			Status:     true,
		}
		if err != nil {
			errs = append(errs, err)
			greConn.Status = false
		}
		greConns = append(greConns, greConn)
	}

	if len(errs) > 0 {
		errStrs := make([]string, 0)
		for _, e := range errs {
			errStrs = append(errStrs, e.Error())
		}
		return greConns, errors.New(strings.Join(errStrs, "\n"))
	}

	return greConns, nil
}

type LibvirtNetwork struct {
	Name       string
	Started    bool
	Autostart  bool
	Persistent bool
}

func RegisterLibvirtNetwork(cl *cluster.Cluster, bridgeName string) (*LibvirtNetwork, error) {
	conn, err := libvirt.NewConnect("qemu:///system") // TODO: system IP
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	xmlString, err := generateNetXML(cl.Name, bridgeName)
	if err != nil {
		return nil, err
	}

	network, err := conn.NetworkDefineXML(xmlString)
	if err != nil {
		return nil, err
	}

	l := &LibvirtNetwork{
		Name:       cl.Name,
		Persistent: true,
	}

	err = network.Create()
	if err != nil {
		return l, err
	}
	l.Started = true

	err = network.SetAutostart(true)
	if err != nil {
		return l, err
	}
	l.Autostart = true

	return l, nil
}
