APP_NAME := shortlink
VERSION := v1.0.0
DB := postgres

build:
	@docker build -t $(APP_NAME):$(VERSION) .

run:
	@docker-compose run -d -e DB=mem -p 8000:8000 --name shortlinkmem shortlink

runsdb:
	@docker-compose up -d

stop:
	@docker stop shortlinkmem

stopsdb:
	@docker stop $(APP_NAME)
	@docker stop $(DB)

log:
	@docker logs $(APP_NAME)

terminal:
	@docker exec -it $(APP_NAME) sh