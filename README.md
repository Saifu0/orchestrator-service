# Orchestrator Service
A simple Orchestrator Service written in Go using gRPC and Protocol Buffer

# Setup and testing locally
- Change directory to your `GOPATH`/src/github.com using command
```bash
cd $GOPATH/src/github.com
```

- Clone this repository


- From the `orchestrator-service` directory, type following command to installing all dependencies
```bash
go mod tidy
``` 

- Open a new terminal and change directory to `datamock` and type following command to installing all dependencies
```bash
go mod tidy
```

- Type below command from `orchestrator-service` directory to run two orchestrator-service instances at port `9000` and `9001`
```bash
go run main
```

- In a new termnial, type below command from `datamock` directory to run datamock server instances at port `10000`
```bash
go run server/server.go
```

- In a new termnial, type below command from `client` directory to send requests to the server
```bash
go run client.go
```

### Note: Above instructions are suitable for Mac and Linux
