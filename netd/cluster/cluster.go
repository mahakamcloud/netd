package cluster

import "fmt"

type Cluster struct {
	Name   string `json:"name"`
	GREKey int    `json:"key"`
}

func New(name string, greKey int) *Cluster {
	return &Cluster{name, greKey}
}

func (c *Cluster) String() string {
	return fmt.Sprintf("Name: %s GREKey: %d", c.Name, c.GREKey)
}
