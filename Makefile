build:
	@echo "Building dev image..."
	@docker-compose build

up:
	@echo "Removing old containers..."
	@docker-compose down
	@echo "Starting containers..."
	@docker-compose up

dev: build up
