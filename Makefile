dev:
	gin --appPort 5000 --port 3000 --build ./cmd --path . --bin ./bin/dev --all

doc:
	swag init --parseDependency --parseInternal -g cmd/main.go -o docs

start:
	go run ./cmd/main.go

build:
	go build -o ./bin/production ./cmd
