package cluster

type Cluster struct {
	name string
	key  int
}

func (c *Cluster) Name() string {
	return c.name
}

func (c *Cluster) Key() int {
	return c.key
}
