package network

import (
	"net"
	"strconv"

	"github.com/digitalocean/go-openvswitch/ovs"
	log "github.com/sirupsen/logrus"
)

func createTapDev(tapDevName string) error {
	ipUtil := NewIPUtil()
	_, err := ipUtil.CreateTapDev(tapDevName)
	return err
}

func createPort(bridgeName, portName string) error {
	o := ovs.New(
		ovs.Sudo(),
	)
	return o.VSwitch.AddPort(bridgeName, portName)
}

// TODO: OVSwitch set interface is always overwritten. Do not execute if interface set.
func createGRETunnel(portName string, remoteIP net.IP, greKey int) error {
	iFaceOptions := ovs.InterfaceOptions{
		Type:     ovs.InterfaceTypeGRE,
		RemoteIP: remoteIP.String(),
		Key:      strconv.Itoa(greKey),
	}
	o := ovs.New(
		ovs.Sudo(),
	)
	return o.VSwitch.Set.Interface(portName, iFaceOptions)
}

func connectToRemote(bridgeName string, tapDevName string, remoteIP net.IP, greKey int) error {
	err := createBridge(bridgeName)
	if err != nil {
		log.Debugf("Error creating bridge %v: %v", bridgeName, err)
		return err
	}

	err = createTapDev(tapDevName)
	if err != nil {
		log.Debugf("Error creating tap device %v: %v", tapDevName, err)
		return err
	}

	err = createPort(bridgeName, tapDevName)
	if err != nil {
		log.Debugf("Error adding port %v to bridge %v: %v", tapDevName, bridgeName, err)
		return err
	}

	err = createGRETunnel(tapDevName, remoteIP, greKey)
	if err != nil {
		log.Debugf("Error creating GRE tunnel on tap device %v with remote IP %v and GRE key %v: %v", tapDevName, remoteIP, greKey, err)
		return err
	}
	return nil
}
