package provisioner

import (
	"encoding/xml"
)

type bridge struct {
	Name  string `xml:"name,attr"`
	Stp   string `xml:"stp,attr"`
	Delay string `xml:"delay,attr"`
}

type virtualport struct {
	PortType string `xml:"type,attr"`
}

type netxml struct {
	XMLName     xml.Name    `xml:"network"`
	Name        string      `xml:"name"`
	Bridge      bridge      `xml:"bridge"`
	Virtualport virtualport `xml:"virtualport"`
}

func generateNetXML(netName, bridgeName string) (string, error) {
	b := bridge{bridgeName, "on", "0"}
	v := virtualport{"openvswitch"}
	n := netxml{
		Name:        netName,
		Bridge:      b,
		Virtualport: v,
	}

	output, err := xml.Marshal(n)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
