package main

import (
	"os"

	"github.com/mahakamcloud/netd/appcontext"
	"github.com/mahakamcloud/netd/config"
	"github.com/mahakamcloud/netd/netd"
	"github.com/mahakamcloud/netd/server"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func handleInitError() {
	if e := recover(); e != nil {
		log.Fatalf("Failed to load the app due to error : %s", e)
	}
}

func main() {
	defer handleInitError()

	err := config.Load()
	if err != nil {
		log.Fatalf("error loading config : %s", err)
	}
	appcontext.Init()

	clientApp := cli.NewApp()
	clientApp.Name = "network-daemon"
	clientApp.Version = "0.0.1"
	clientApp.Commands = []cli.Command{
		{
			Name:        "start",
			Description: "Start HTTP api server",
			Action: func(c *cli.Context) error {
				if err := netd.Register(); err != nil {
					log.Errorf("failed host registration: %v", err)
					return err
				}

				server.StartAPIServer()
				return nil
			},
		},
	}
	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
