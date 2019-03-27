package provisioner

import (
	"encoding/xml"
)

type networkInterface struct {
	IfaceType   string        `xml:"type,attr"`
	Source      networkSource `xml:"source"`
	Virtualport virtualport   `xml:"virtualport"`
}

type networkSource struct {
	BridgeName string `xml:"bridge,attr"`
}

type virtualport struct {
	PortType string `xml:"type,attr"`
}

type netxml struct {
	XMLName          xml.Name         `xml:"network"`
	Name             string           `xml:"name"`
	NetworkInterface networkInterface `xml:"interface"`
}

func generateNetXML(netName, bridgeName string) (string, error) {
	i := networkInterface{
		IfaceType:   "bridge",
		Source:      networkSource{bridgeName},
		Virtualport: virtualport{"openvswitch"},
	}
	n := netxml{
		Name:             netName,
		NetworkInterface: i,
	}

	output, err := xml.Marshal(n)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
