build:
	protoc --proto_path=$(GOPATH)/src/github.com/hoflish/shippy/services/consignment \
	--micro_out=. \
	--go_out=. proto/consignment.proto

	docker build -t consignment-service .

run: 
	docker run -p 50051:50051 consignment-service
