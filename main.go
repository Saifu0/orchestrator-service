package main

import (
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Saifu0/orchestrator-service/logic"
	protos "github.com/Saifu0/orchestrator-service/protos/orchestrator"
)

func createServer(port int) {
	// Initializing a logger for logging error messgaes
	log := hclog.Default()

	gs := grpc.NewServer()
	oc := logic.NewOrchestrator(log)

	protos.RegisterOrchestratorServer(gs, oc)
	reflection.Register(gs)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Error("Unable to create listener at ", port, " error", err)
		os.Exit(1)
	}

	gs.Serve(listener)
}

func main() {
	// create a WaitGroup
	wg := new(sync.WaitGroup)

	// add two goroutines to 'wait' WaitGroup
	wg.Add(2)

	// goroutine to launch a server on port 9000
	go createServer(9000)

	// goroutine to launch a server on port 9001
	go createServer(9001)

	wg.Wait()
}
