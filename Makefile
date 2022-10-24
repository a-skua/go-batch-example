upload:
	docker compose up

fmt:
	docker compose run --rm gopher go fmt ./...
