dev:
	gin --appPort 5000 --port 3000 run cmd/main.go

start:
	go run cmd/main.go

build:
	go build -o bin/app cmd/main.go