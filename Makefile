.PHONY: build up down remove help

build: ## Build docker image
	docker-compose build --progress=plain --no-cache
up: ## Do docker compose up
	docker-compose up -d
down: ## Do docker compose down
	docker-compose down
remove: ## Do 'make down' and remove containers for services not defined in the Compose file
	docker-compose down --remove-orphans

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
