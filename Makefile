build:
	docker compose up --build -d

lint:
	golangci-lint run -c .golangci.yml

server: lint
	@bash scripts/server.sh

stop:
	docker compose kill

log:
	docker logs ledger
