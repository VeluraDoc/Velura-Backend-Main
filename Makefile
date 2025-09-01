SWAG_MAIN=cmd/main.go
SWAG_OUT=docs

.PHONY: doc dev start build
doc:
	swag init -g ./$(SWAG_MAIN) -o ./$(SWAG_OUT) --parseDependency --parseInternal

dev:
	gin --appPort 5000 --port 3000 --build ./cmd --path . --bin ./bin/dev --all

start:
	go run ./cmd/main.go

build:
	go build -o ./bin/production ./cmd

test:
	go test ./...
