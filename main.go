package main

import (
	"log"
	"time"

	"github.com/selftechio/sandhog/internal/tunnel"
)

func main() {
	tun, err := tunnel.NewTunnel("10.11.12.13/32", "wgtest0")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer tun.Close()

	time.Sleep(time.Second * 5)
}
