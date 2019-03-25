package network

import (
	"errors"
	"fmt"

	"github.com/digitalocean/go-openvswitch/ovs"
)

type Bridge struct {
	name string
}

func NewBridge(name string) (*Bridge, error) {
	b := &Bridge{name}
	o := ovs.New()
	err := o.VSwitch.AddBridge(name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error creating a new bridge %v: %v", name, err))
	}
	return b, nil
}

func (b *Bridge) Name() string {
	return b.name
}
