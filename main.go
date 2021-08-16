package main

import (
	"net"
	"os"

	protos "github.com/Saifu0/orchestrator-service/protos/orchestrator"
	"github.com/Saifu0/orchestrator-service/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	oc := server.NewOrchestrator(log)

	protos.RegisterOrchestratorServer(gs, oc)
	reflection.Register(gs)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	gs.Serve(listener)
}
