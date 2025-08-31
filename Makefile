dev:
	gin --appPort 5000 --port 3000 --build ./cmd --path . --bin ./bin/dev --all

start:
	go run ./cmd/main.go

build:
	go build -o ./bin/production ./cmd
