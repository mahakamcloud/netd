package server

import (
	"github.com/gorilla/mux"
	"github.com/mahakamcloud/netd/service"
)

func Router(services *service.Services) *mux.Router {
	router := mux.NewRouter()
	return router
}
