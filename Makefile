# Makefile
.PHONY: test docker-dev-up docker-dev-down docker-prod-up docker-dev-down

# Определение переменных
DOCKER_COMPOSE_DEV = ./docker-compose/docker-compose.dev.yml
DOCKER_COMPOSE_PROD = ./docker-compose/docker-compose.prod.yml

run:
	go run ./cmd/app/main.go
# Команды для запуска Docker Compose в режиме разработки
# d - docker, d - dev, b - build, u - up, d - down
ddb:
	docker-compose -f $(DOCKER_COMPOSE_DEV) build
	
ddu:
	docker-compose -f $(DOCKER_COMPOSE_DEV) up -d

ddd:
	docker-compose -f $(DOCKER_COMPOSE_DEV) down

# Команды для запуска Docker Compose в режиме продакшн
# d - docker, p - prod, b - build, u - up, d - down
dpb:
	docker-compose -f $(DOCKER_COMPOSE_PROD) build
	
dpu:
	docker-compose -f $(DOCKER_COMPOSE_PROD) up -d

dpd:
	docker-compose -f $(DOCKER_COMPOSE_PROD) down

