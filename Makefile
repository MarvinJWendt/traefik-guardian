build-image:
	@echo "Building dev image..."
	@docker-compose build

up:
	@echo "Removing old containers..."
	@docker-compose down
	@echo "Starting containers..."
	@docker-compose up

format:
	@echo Formatting Go code...
	@(cd src && go fmt ./...)
	@echo Formatting frontend code...
	@(cd src/admin-ui && npm run format)

build: format build-image
dev: build up
