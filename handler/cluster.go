package handler

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/mahakamcloud/netd/logger"
	"github.com/mahakamcloud/netd/netd/cluster"
	"github.com/mahakamcloud/netd/netd/host"
	"github.com/mahakamcloud/netd/netd/provisioner"
	log "github.com/sirupsen/logrus"
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

	response.BridgeResp = createBridge(req.Cluster)
	if !response.BridgeResp.Status {
		response.Status = false
		w.WriteHeader(http.StatusInternalServerError)
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		log.Debug(string(responseJSON))
		return
	}

	bridgeName := response.BridgeResp.Name

	tunnelResp, tunnelStatus := createGRETunnels(req.Cluster, req.Hosts, bridgeName)
	response.GRETunnelsResp = tunnelResp
	response.LibvirtNetResp = createLibvirtNetwork(req.Cluster, bridgeName)

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

func createBridge(cl *cluster.Cluster) *bridgeResp {
	bridge, err := provisioner.CreateBridge(cl)
	if err != nil {
		return &bridgeResp{
			Status: false,
			Err:    err.Error(),
		}
	}

	return &bridgeResp{
		Name:   bridge.Name(),
		Status: true,
	}
}

func createGRETunnels(cl *cluster.Cluster, hosts []*host.Host, bridgeName string) ([]*greTunnelResp, bool) {
	//TODO Get appropriate localhost
	localhost := host.New("local", net.IPv4(10, 0, 2, 15), net.IPv4Mask(255, 255, 255, 0))

	status := true
	greConns, errs := provisioner.CreateGREMesh(cl, localhost, hosts, bridgeName)
	if errs != nil {
		status = false
	}

	greTuRe := make([]*greTunnelResp, 0)
	for _, g := range greConns {
		greTuRe = append(greTuRe, &greTunnelResp{g.Name, g.RemoteHost, g.Status, errs.Error()})
	}
	return greTuRe, status
}

func createLibvirtNetwork(cl *cluster.Cluster, bridgeName string) *libvirtNetResp {
	libvirtNet, err := provisioner.RegisterLibvirtNetwork(cl, bridgeName)
	if err != nil {
		return &libvirtNetResp{
			Status: false,
			Err:    err.Error(),
		}
	}

	return &libvirtNetResp{
		Name:   libvirtNet.Name,
		Status: libvirtNet.Persistent && libvirtNet.Started && libvirtNet.Autostart,
	}
}
