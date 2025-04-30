.PHONY: build
build:
	docker buildx bake runtime

.PHONY: up
up:
	docker compose up

.PHONY: down
down:
	docker compose down

.PHONY: exec
exec:
	docker exec -it runtime /bin/sh
