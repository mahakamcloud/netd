package cluster

type Cluster struct {
	Name   string `json:"name"`
	GREKey int    `json:"key"`
}

func New(name string, greKey int) *Cluster {
	return &Cluster{name, greKey}
}
