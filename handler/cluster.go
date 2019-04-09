package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mahakamcloud/netd/config"
	"github.com/mahakamcloud/netd/logger"
	"github.com/mahakamcloud/netd/netd/cluster"
	"github.com/mahakamcloud/netd/netd/host"
	"github.com/mahakamcloud/netd/netd/network"
	"github.com/mahakamcloud/netd/netd/provisioner"
	log "github.com/sirupsen/logrus"
)

type createClusterNetworkRequest struct {
	Cluster *cluster.Cluster `json:"cluster"`
	Hosts   []*host.Host     `json:"hosts"`
}

type libvirtNetResp struct {
	Name       string `json:"name"`
	BridgeName string `json:"br_name"`
	Status     bool   `json:"status"`
	Err        string `json:"error"`
}

type greTunnelResp struct {
	Name   string     `json:"name"`
	Host   *host.Host `json:"host"`
	Status bool       `json:"status"`
	Err    string     `json:"error"`
}

type createClusterNetworkResponse struct {
	Status         bool             `json:"status"`
	LibvirtNetResp *libvirtNetResp  `json:"libvirtnet"`
	GRETunnelsResp []*greTunnelResp `json:"gre_tunnels"`
}

func CreateClusterNetworkHandler(w http.ResponseWriter, r *http.Request) {
	var req createClusterNetworkRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Log.Error(err)
		return
	}

	if req.Cluster == nil || len(req.Hosts) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		logger.Log.Error(err)
		return
	}

	response := &createClusterNetworkResponse{}

	response.LibvirtNetResp = createBridgeWithLibvirtNetwork(req.Cluster)
	if !response.LibvirtNetResp.Status {
		response.Status = false
		w.WriteHeader(http.StatusInternalServerError)
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		log.Debug(string(responseJSON))
		return
	}

	bridgeName := response.LibvirtNetResp.BridgeName

	localIP, localIPMask, err := network.GetBridgeIPNet(config.HostBridgeName())
	if err != nil {
		response.Status = false
		w.WriteHeader(http.StatusInternalServerError)
		logger.Log.Error(err)
		return
	}

	localhost := host.New("local", localIP, localIPMask)
	tunnelResp, tunnelStatus := createGRETunnels(req.Cluster, localhost, req.Hosts, bridgeName)
	response.GRETunnelsResp = tunnelResp

	if !tunnelStatus || !response.LibvirtNetResp.Status {
		response.Status = false
		w.WriteHeader(http.StatusMultiStatus)
	} else {
		response.Status = true
		w.WriteHeader(http.StatusCreated)
	}

	responseJSON, _ := json.Marshal(response)
	w.Write(responseJSON)
	logger.Log.Debug(responseJSON)
}

func createBridgeWithLibvirtNetwork(cl *cluster.Cluster) *libvirtNetResp {
	libvirtnet, err := provisioner.CreateBridgeWithLibvirtNetwork(cl)
	if err != nil {
		return &libvirtNetResp{
			Status: false,
			Err:    err.Error(),
		}
	}

	return &libvirtNetResp{
		Name:       libvirtnet.Name,
		BridgeName: libvirtnet.BridgeName,
		Status:     libvirtnet.Persistent && libvirtnet.Started && libvirtnet.Autostart,
	}
}

func createGRETunnels(cl *cluster.Cluster, localhost *host.Host, hosts []*host.Host, bridgeName string) ([]*greTunnelResp, bool) {
	status := true
	greConns, err := provisioner.CreateGREMesh(cl, localhost, hosts, bridgeName)
	if err != nil {
		status = false
	}

	greTunResp := make([]*greTunnelResp, 0)
	for _, g := range greConns {
		r := &greTunnelResp{
			Name:   g.Name,
			Host:   g.RemoteHost,
			Status: g.Status,
		}
		if err != nil {
			r.Err = err.Error()
		}
		greTunResp = append(greTunResp, r)
	}
	return greTunResp, status
}
