build:
	@go build -o bin/api
run: build
	@./bin/api
test:
	@go test -v ./...
wire:
	@go run github.com/google/wire/cmd/wire cmd/server/