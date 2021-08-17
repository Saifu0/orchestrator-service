package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	protos "github.com/Saifu0/orchestrator-service/protos/orchestrator"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:9000", "The server address in the format of host:port")
)

func main() {
	log := hclog.Default()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Error("Unable to make connection with server!", "error", err)
		os.Exit(1)
	}

	client := protos.NewOrchestratorClient(conn)

	resp, err := client.GetUserByName(context.Background(), &protos.Request{Name: "Saifur"})
	if err != nil {
		log.Error("Unable to get the data ", err)
		os.Exit(1)
	}

	fmt.Println("[RESPONSE]: ", resp)
}
