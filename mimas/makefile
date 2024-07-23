.PHONY: default run

APP_NAME = mimas

default: run

run:
	@docker-compose up -d
build:
	go build main.go
test:
	go test ./...
docs:
	@swag init 
protoc:
	@protoc --go_out=. --go-grpc_out=. ./proto/product.proto        
clean:
	@rm -rf docs