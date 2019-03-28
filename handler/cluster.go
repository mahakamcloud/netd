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
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Cluster == nil || len(req.Hosts) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	response := &createClusterNetworkResponse{}
	response.BridgeResp = createBridge(req.Cluster)

	bridgeName := response.BridgeResp.Name

	response.GRETunnelsResp = createGRETunnels(req.Cluster, req.Hosts, bridgeName)
	response.LibvirtNetResp = createLibvirtNetwork(req.Cluster, bridgeName)
	response.Status = true

	responseJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}

func createBridge(cl *cluster.Cluster) *bridgeResp {
	bridge, err := provisioner.CreateBridge(cl)
	if err != nil {
		return nil
	}

	return &bridgeResp{
		Name:   bridge.Name(),
		Status: true,
	}
}

func createGRETunnels(cl *cluster.Cluster, hosts []*host.Host, bridgeName string) []*greTunnelResp {
	//TODO Get appropriate localhost
	localhost := host.New("local", net.IPv4(10, 0, 2, 15), net.IPv4Mask(255, 255, 255, 0))
	greConns, errs := provisioner.CreateGREMesh(cl, localhost, hosts, bridgeName)
	if errs != nil {
		fmt.Println(errs)
	}

	greTuRe := make([]*greTunnelResp, 0)
	for _, g := range greConns {
		greTuRe = append(greTuRe, &greTunnelResp{g.Name, g.RemoteHost, g.Status, ""})
	}
	return greTuRe
}

func createLibvirtNetwork(cl *cluster.Cluster, bridgeName string) *libvirtNetResp {
	libvirtNet, err := provisioner.RegisterLibvirtNetwork(cl, bridgeName)
	if err != nil {
		return nil
	}
	return &libvirtNetResp{
		Name:   libvirtNet.Name,
		Status: libvirtNet.Persistent && libvirtNet.Started && libvirtNet.Autostart,
	}
}
