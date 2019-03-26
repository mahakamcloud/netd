package cluster

type Cluster struct {
	name   string `json:"name"`
	greKey int    `json:"key"`
}

func New(name string, greKey int) *Cluster {
	return &Cluster{name, greKey}
}

func (c *Cluster) Name() string {
	return c.name
}

func (c *Cluster) GREKey() int {
	return c.greKey
}
