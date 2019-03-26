package main

import (
	"github.com/mahakamcloud/netd/appcontext"
	"github.com/mahakamcloud/netd/config"
	log "github.com/sirupsen/logrus"
)

func handleInitError() {
	if e := recover(); e != nil {
		log.Fatalf("Failed to load the app due to error : %s", e)
	}
}

func main() {
	defer handleInitError()

	config.Load()
	appcontext.Init()
}
