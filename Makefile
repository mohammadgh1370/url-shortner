# Makefile

build:
	@cp -n .env.example .env
	@docker-compose up --build -d

up:
	@docker-compose up -d

down:
	@docker-compose down

exec-app:
	@docker exec -it app sh

test:
	@docker exec -i app sh -c "go test -v ./internal/controller/"
