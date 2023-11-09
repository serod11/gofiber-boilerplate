run:
	go run cmd/server/main.go

run-live:
	air

build:
	go build -o build/main cmd/server/main.go