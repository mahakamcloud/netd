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

func ProvisionClusterNetwork(cl *cluster.Cluster, localhost *host.Host, remotehosts []*host.Host) error {
	bridgeName := generateBridgeName(cl.Name)
	bridge, err := network.NewBridge(bridgeName)
	if err != nil {
		return err
	}

	// create GRE mesh
	// TODO: do goroutines
	// TODO: make error messages Mahakam way
	errs := make([]error, 0)
	for _, r := range remotehosts {
		greName := generateGRETunnelName(cl.Name, localhost.Name(), r.Name())
		gre := network.NewGRE(greName, r.IPAddr(), cl.GREKey)
		err = gre.Create(bridgeName)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		errStrs := make([]string, 0)
		for _, e := range errs {
			errStrs = append(errStrs, e.Error())
		}
		return errors.New(strings.Join(errStrs, "\n"))
	}
	// register libvirt network

	err = registerLibvirtNetwork(cl.Name, bridge)
	if err != nil {
		return err
	}

	return nil
}

func registerLibvirtNetwork(name string, br *network.Bridge) error {
	conn, err := libvirt.NewConnect("qemu:///system") // TODO: system IP
	if err != nil {
		return err
	}
	defer conn.Close()

	xmlString, err := generateNetXML(name, br.Name())
	if err != nil {
		return err
	}

	network, err := conn.NetworkDefineXML(xmlString)
	if err != nil {
		return err
	}

	err = network.Create()
	if err != nil {
		return err
	}

	err = network.SetAutostart(true)
	if err != nil {
		return err
	}

	return nil
}
