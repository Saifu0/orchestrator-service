package main

import (
	"context"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	protos "github.com/Saifu0/orchestrator-service/protos/orchestrator"
)

type Orchestrator struct {
	log hclog.Logger
}

func NewOrchestrator(l hclog.Logger) *Orchestrator {
	oc := &Orchestrator{log: l}
	return oc
}

func (oc *Orchestrator) GetUserByName(ctx context.Context, r *protos.Request) (*protos.Response, error) {
	oc.log.Info("Handle request for GetUserByName, name: ", r.GetName())

	return &protos.Response{Name: r.GetName(), Class: "CS", Roll: 176231}, nil
}

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	oc := NewOrchestrator(log)

	protos.RegisterOrchestratorServer(gs, oc)
	reflection.Register(gs)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	gs.Serve(listener)
}
