# Makefile

build:
	@cp -n .env.example .env
	@docker-compose up --build -d

up:
	@docker-compose up -d

down:
	@docker-compose down -v
