package provisioner

import (
	"encoding/xml"

	"github.com/mahakamcloud/netd/logger"
)

type bridge struct {
	Name  string `xml:"name,attr"`
	STP   string `xml:"stp,attr"`
	Delay string `xml:"delay,attr"`
}

type netxml struct {
	XMLName xml.Name `xml:"network"`
	Name    string   `xml:"name"`
	Bridge  bridge   `xml:"bridge"`
}

func generateNetXML(netName, bridgeName string) (string, error) {
	b := bridge{bridgeName, "on", "0"}
	n := netxml{
		Name:   netName,
		Bridge: b,
	}

	output, err := xml.Marshal(n)
	if err != nil {
		logger.Log.Error(err)
		return "", err
	}
	return string(output), nil
}
