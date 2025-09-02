SWAG_MAIN=cmd/api/main.go
SWAG_OUT=docs

.PHONY: doc dev start build
doc:
	swag init -g ./$(SWAG_MAIN) -o ./$(SWAG_OUT) --parseDependency --parseInternal

start-cli:
	go run ./cmd/cli/main.go $(ARGS)

start-api:
	go run ./cmd/api/main.go

build:
	go build -o ./bin/production ./cmd

test:
	go test ./...
