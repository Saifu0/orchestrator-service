package logic

import (
	"context"

	protos "github.com/Saifu0/orchestrator-service/protos/orchestrator"
	"github.com/hashicorp/go-hclog"
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
