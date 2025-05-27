run:
	go run cmd/app/main.go
	
dev:
	docker compose up --build

wire:
	wire ./internal/wire/wire.go

run-blockchain:
	cd blockchain && dfx deploy