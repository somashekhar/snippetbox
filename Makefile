.PHONY: help db-up db-down db-restart db-logs db-shell run

help:
	@echo "Available commands:"
	@echo "  make db-up       - Start MySQL container"
	@echo "  make db-down     - Stop MySQL container"
	@echo "  make db-restart  - Restart MySQL container"
	@echo "  make db-logs     - View MySQL logs"
	@echo "  make db-shell    - Connect to MySQL shell"
	@echo "  make run         - Run the Go application"

db-up:
	docker-compose up -d

db-down:
	docker-compose down

db-restart:
	docker-compose restart

db-logs:
	docker-compose logs -f mysql

db-shell:
	docker exec -it snippetbox_mysql mysql -u web -ppass snippetbox

run:
	go run ./cmd/web
