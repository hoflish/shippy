build:
	protoc --proto_path=$(GOPATH)/src/github.com/hoflish/shippy/services/vessel \
		--micro_out=. \
		--go_out=. proto/vessel.proto
	
	docker build -t vessel-service .

run:
	docker run -p 50052:50051 \
	-e MICRO_SERVER_ADDRESS=:50051 \
	-e MICRO_REGISTRY=mdns vessel-service
