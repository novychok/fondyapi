build-client:
	@go build -o ./bin/client ./client/client.go

build-service:
	@go build -o ./bin/service ./service/service.go

client: build-client
	@./bin/client

service: build-service
	@./bin/service