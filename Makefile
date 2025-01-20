build:
	docker compose up --build -d

server:
	@bash scripts/server.sh

stop:
	docker compose kill

log:
	docker logs ledger
