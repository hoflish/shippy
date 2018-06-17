build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/hoflish/shippy/services/consignment \
	proto/consignment.proto

	docker build -t consignment-service .

run: 
	docker run -p 50051:50051 consignment-service
