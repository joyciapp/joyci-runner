package main

import (
	"log"
	"time"

	"github.com/joyciapp/joyci-grpc/grpc/api"
)

func main() {
	log.Println("Run git clone...")
	api.GitClone("git@github.com:joyciapp/joyci-runner.git")

	time.Sleep(30 * time.Second)
}
