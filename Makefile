# Makefile

.PHONY: docker-build docker-run docker-stop docker-clean docker-logs

docker-build:
	@echo "Building Docker image..."
	@docker-compose build

docker-run:
	@echo "Running Docker containers..."
	@docker-compose up -d

docker-stop:
	@echo "Stopping Docker containers..."
	@docker-compose down

docker-clean:
	@echo "Cleaning up Docker resources..."
	@docker-compose down --volumes

docker-logs:
	@echo "Viewing Docker logs..."
	@docker-compose logs -f
