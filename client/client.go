package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	protos "github.com/Saifu0/orchestrator-service/simple-rpc/protos/orchestrator"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:8080", "The server address in the format of host:port")
)

func main() {
	log := hclog.Default()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Error("Unable in connection with server!", "error", err)
		os.Exit(1)
	}

	client := protos.NewOrchestratorClient(conn)

	resp, err := client.GetUserByName(context.Background(), &protos.Request{Name: "Saifurrahman"})
	if err != nil {
		log.Error("Unable to call service method, GetUserByName, ", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}
