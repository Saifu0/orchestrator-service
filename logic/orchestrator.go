package logic

import (
	"context"
	"fmt"

	datamock "github.com/Saifu0/orchestrator-service/datamock/protos/datamock"
	protos "github.com/Saifu0/orchestrator-service/protos/orchestrator"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

type Orchestrator struct {
	log hclog.Logger
}

func NewOrchestrator(l hclog.Logger) *Orchestrator {
	oc := &Orchestrator{log: l}
	return oc
}

func makeClient(serverAddr string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Errorf("unable to make connection with server, %s", serverAddr)
		return nil, err
	}

	return conn, nil
}

func (oc *Orchestrator) GetUserByName(ctx context.Context, r *protos.Request) (*protos.Response, error) {
	oc.log.Info("Handle request for GetUserByName method, name: %s", r.GetName())

	// making a connection to port=9001 and calling GetUser method
	conn, err := makeClient("localhost:9001")
	if err != nil {
		return nil, err
	}
	client := protos.NewOrchestratorClient(conn)
	resp, err := client.GetUser(context.Background(), &protos.Request{Name: r.GetName()})

	return resp, nil
}

func (oc *Orchestrator) GetUser(ctx context.Context, r *protos.Request) (*protos.Response, error) {
	oc.log.Info("Handle request for GetUser method, name: ", r.GetName())

	// making a connection to port=10000 and calling GetMockUserData method
	conn, err := makeClient("localhost:10000")
	if err != nil {
		return nil, err
	}
	client := datamock.NewDatamockClient(conn)
	resp, err := client.GetMockUserData(context.Background(), &datamock.Request{Name: r.GetName()})

	if err != nil {
		return nil, err
	}
	// serializing data into datamock.Response format
	orchestrator_response := &protos.Response{
		Name:  resp.Name,
		Class: resp.Class,
		Roll:  resp.Roll,
	}

	return orchestrator_response, nil
}

// func redirectRequest(name string) ()
