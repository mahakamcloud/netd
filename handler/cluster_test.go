package handler

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mahakamcloud/netd/config"
	"github.com/mahakamcloud/netd/logger"
	"github.com/mahakamcloud/netd/netd/cluster"
	"github.com/mahakamcloud/netd/netd/host"
	"github.com/stretchr/testify/assert"
)

func setup() {
	config.Load()
	logger.SetupLogger()
}

func TestCreateClusterNetworkHandler(t *testing.T) {
	setup()

	cl1 := cluster.New("cl1", 1)
	h1 := host.New("host1", net.IPv4(10, 10, 10, 1), net.IPv4Mask(255, 255, 255, 0))
	h2 := host.New("host2", net.IPv4(10, 10, 10, 2), net.IPv4Mask(255, 255, 255, 0))
	req := createClusterNetworkRequest{cl1, []*host.Host{h1, h2}}
	reqJSON, _ := json.Marshal(req)

	b := &bridgeResp{"cl1_br", true, ""}
	g1 := &greTunnelResp{"cl1_local_host1", h1, true, ""}
	g2 := &greTunnelResp{"cl1_local_host2", h2, true, ""}
	l := &libvirtNetResp{"cl1", true, ""}
	response := &createClusterNetworkResponse{true, b, []*greTunnelResp{g1, g2}, l}
	responseJSON, _ := json.Marshal(response)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/v1/network", bytes.NewBuffer(reqJSON))

	CreateClusterNetworkHandler(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, string(responseJSON), w.Body.String())
}

func TestShouldReturn422ForUnprocessableJSON(t *testing.T) {
	setup()

	invalidReqJSON := "{\"foo\":\"bar\"}"

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/v1/network", bytes.NewBufferString(invalidReqJSON))

	CreateClusterNetworkHandler(w, r)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestShouldReturn400ForInvalidJSON(t *testing.T) {
	setup()

	invalidReqJSON := "{\"foo\":\"bar\""

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/v1/network", bytes.NewBufferString(invalidReqJSON))

	CreateClusterNetworkHandler(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestCreateLibvirtNetworkForProvisionerError tests if libvirt network creation fails.
//
// One sureshot way of ensuring that libvirt network creation fails is by trying to create the same network
// twice. That is why this test sends the same request twice.
// TODO: Find a better way of doing this.
func TestCreateLibvirtNetworkForProvisionerError(t *testing.T) {
	setup()

	cl1 := cluster.New("cl2", 2)
	h1 := host.New("host3", net.IPv4(10, 10, 10, 3), net.IPv4Mask(255, 255, 255, 0))
	h2 := host.New("host4", net.IPv4(10, 10, 10, 4), net.IPv4Mask(255, 255, 255, 0))
	req := createClusterNetworkRequest{cl1, []*host.Host{h1, h2}}
	reqJSON, _ := json.Marshal(req)

	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	r1, _ := http.NewRequest("POST", "/v1/network", bytes.NewBuffer(reqJSON))
	r2, _ := http.NewRequest("POST", "/v1/network", bytes.NewBuffer(reqJSON))

	CreateClusterNetworkHandler(w1, r1)
	CreateClusterNetworkHandler(w2, r2)

	assert.Equal(t, http.StatusMultiStatus, w2.Code)
	assert.Equal(t, "application/json", w2.Header().Get("Content-Type"))
	assert.Contains(t, w2.Body.String(), "operation failed: network 'cl2' already exists with uuid")
}
