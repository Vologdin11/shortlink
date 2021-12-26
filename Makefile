APP_NAME := shortlink
VERSION := v1.0.0

build:
	@docker build -t $(APP_NAME):$(VERSION) .

run:
	@docker-compose up -d

postgres:
	@docker run --name postgres -e POSTGRES_PASSWORD=qweasd -p 5432:5432 -d --rm postgres

stop:
	@docker stop $(APP_NAME)

log:
	@docker logs $(APP_NAME)

terminal:
	@docker exec -it $(APP_NAME) sh