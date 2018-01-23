run:
	@go run ./cmd/server/main.go

test:
	@go test ./pkg/...

up:
	@docker-compose up --build
