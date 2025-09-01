SWAG_MAIN=cmd/main.go
SWAG_OUT=docs

.PHONY: doc
doc:
	swag init --parseDependency --parseInternal \
	  -g $(SWAG_MAIN) \
	  --dir cmd \
	  --dir internal \
	  -o $(SWAG_OUT)

dev:
	gin --appPort 5000 --port 3000 --build ./cmd --path . --bin ./bin/dev --all

start:
	go run ./cmd/main.go

build:
	go build -o ./bin/production ./cmd
