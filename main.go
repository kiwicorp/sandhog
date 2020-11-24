package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/go-hclog"
	"github.com/selftechio/sandhog/internal/server"
)

var (
	log hclog.Logger
)

func init() {
	log = hclog.Default().Named("main")
}

func main() {
	log.Info("hello, this is sandhog")

	server.Start()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan
}
