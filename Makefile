migrate:
	@go run ./cmd/migrate/main.go --config=./config/local.yaml --migrations-path=./migrations

build:
	@mkdir -p ./bin
	@go build -o ./bin/image_service ./cmd/image_service/main.go --config=./config/local.yaml

run: build
	@./bin/image_service

test:
	@go test -v ./...

proto:
	@command -v protoc >/dev/null 2>&1 || { echo "protoc is required but not installed. Aborting."; exit 1; }
	@protoc -I proto proto/image/image_service.proto --go_out=./pkg/gen/go/ --go_opt=paths=source_relative --go-grpc_out=./pkg/gen/go/ --go-grpc_opt=paths=source_relative
