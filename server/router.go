package server

import (
	"github.com/gorilla/mux"
	"github.com/mahakamcloud/netd/handler"
	"github.com/mahakamcloud/netd/service"
)

func Router(services *service.Services) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	router.HandleFunc("/v1/network", handler.CreateClusterNetworkHandler).Methods("POST")
	return router
}
