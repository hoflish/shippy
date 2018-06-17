build:
	protoc --proto_path=$(GOPATH)/src/github.com/hoflish/shippy/services/consignment \
		--micro_out=. \
		--go_out=. proto/consignment.proto

	docker build -t consignment-service .

run: 
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns consignment-service
