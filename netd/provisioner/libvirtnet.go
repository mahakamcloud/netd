package provisioner

import (
	"encoding/xml"

	"github.com/mahakamcloud/netd/logger"
)

type forward struct {
	Mode string `xml:"mode,attr"`
}

type bridge struct {
	Name string `xml:"name,attr"`
}

type virtualport struct {
	PortType string `xml:"type,attr"`
}

type netxml struct {
	XMLName     xml.Name    `xml:"network"`
	Name        string      `xml:"name"`
	Forward     forward     `xml:"forward"`
	Bridge      bridge      `xml:"bridge"`
	Virtualport virtualport `xml:"virtualport"`
}

func generateNetXML(netName, bridgeName string) (string, error) {
	f := forward{"bridge"}
	b := bridge{bridgeName}
	v := virtualport{"openvswitch"}
	n := netxml{
		Name:        netName,
		Forward:     f,
		Bridge:      b,
		Virtualport: v,
	}

	output, err := xml.Marshal(n)
	if err != nil {
		logger.Log.Error(err)
		return "", err
	}
	return string(output), nil
}
