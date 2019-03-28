package handler

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/mahakamcloud/netd/netd/cluster"
	"github.com/mahakamcloud/netd/netd/host"
	"github.com/mahakamcloud/netd/netd/provisioner"
)

type createClusterNetworkRequest struct {
	Cluster *cluster.Cluster `json:"cluster"`
	Hosts   []*host.Host     `json:"hosts"`
}

type bridgeResp struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
	Err    string `json:"error"`
}

type greTunnelResp struct {
	Name   string     `json:"name"`
	Host   *host.Host `json:"host"`
	Status bool       `json:"status"`
	Err    string     `json:"error"`
}

type libvirtNetResp struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
	Err    string `json:"error"`
}

type createClusterNetworkResponse struct {
	Status         bool             `json:"status"`
	BridgeResp     *bridgeResp      `json:"bridge"`
	GRETunnelsResp []*greTunnelResp `json:"gre_tunnels"`
	LibvirtNetResp *libvirtNetResp  `json:"libvirtnet"`
}

func CreateClusterNetworkHandler(w http.ResponseWriter, r *http.Request) {
	var req createClusterNetworkRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	fmt.Println(err)
	defer r.Body.Close()

	response := &createClusterNetworkResponse{}

	brRe := &bridgeResp{}
	bridge, err := provisioner.CreateBridge(req.Cluster)
	if err != nil {
		fmt.Println(err)
	}
	brRe.Name = bridge.Name()
	brRe.Status = true
	response.BridgeResp = brRe

	greTuRe := make([]*greTunnelResp, 0)
	localhost := host.New("local", net.IPv4(10, 0, 2, 15), net.IPv4Mask(255, 255, 255, 0))
	greConns, errs := provisioner.CreateGREMesh(req.Cluster, localhost, req.Hosts, bridge)
	if errs != nil {
		fmt.Println(err)
	}
	for _, g := range greConns {
		greTuRe = append(greTuRe, &greTunnelResp{g.Name, g.RemoteHost, g.Status, ""})
	}
	response.GRETunnelsResp = greTuRe

	liNetRs := &libvirtNetResp{}
	libvirtNet, err := provisioner.RegisterLibvirtNetwork(req.Cluster, bridge)
	if err != nil {
		fmt.Println(err)
	}
	liNetRs.Name = libvirtNet.Name
	liNetRs.Status = libvirtNet.Persistent && libvirtNet.Started && libvirtNet.Autostart
	liNetRs.Err = ""
	response.LibvirtNetResp = liNetRs

	response.Status = true

	responseJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}
