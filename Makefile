.PHONY: protos

protos:
	 protoc -I protos/ protos/orchestrator.proto --go_out=plugins=grpc:protos/orchestrator