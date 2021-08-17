package main

import (
	"context"
	"errors"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	protos "github.com/Saifu0/orchestrator-service/datamock/protos/datamock"
)

type Datamock struct {
	log hclog.Logger
}

func NewDatamock(l hclog.Logger) *Datamock {
	oc := &Datamock{log: l}
	return oc
}

func (d *Datamock) GetMockUserData(ctx context.Context, r *protos.Request) (*protos.Response, error) {
	d.log.Info("Handle request for GetMockUserData, name: ", r.GetName())

	if len(r.GetName()) > 6 {
		return nil, errors.New("length of name is exceeded, must be less than 6")
	}
	length := len(r.GetName())
	return &protos.Response{Name: r.GetName(), Class: int64(length), Roll: int64(length * 10)}, nil
}

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	d := NewDatamock(log)

	protos.RegisterDatamockServer(gs, d)
	reflection.Register(gs)

	listener, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	gs.Serve(listener)
}
