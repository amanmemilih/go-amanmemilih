run:
	fresh
	
dev:
	docker compose up --build

wire:
	wire ./internal/wire/wire.go

run-blockchain:
	cd blockchain && dfx deploy