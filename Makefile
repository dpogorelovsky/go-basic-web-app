install:
	go mod download

run:
	go run cmd/server/main.go

migrate\:up:
	go run cmd/migrate/main.go -direction up

migrate\:down:
	go run cmd/migrate/main.go -direction down