build:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o app cmd/cli/main.go

docker-build:
	docker-compose -f docker-compose.dev.yml build

enter-container:
	docker-compose -f docker-compose.dev.yml exec app bash

run-db:
	docker-compose -f docker-compose.dev.yml up -d db

run-app:
	docker-compose -f docker-compose.dev.yml up -d --force-recreate

build-run: docker-build
	$(MAKE) run-app
	$(MAKE) enter-container