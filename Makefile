build:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o app cmd/cli/main.go

run-db:
	docker-compose up db -d

compose-build:
	docker-compose build

exec-bash:
	docker-compose exec app bash

run-dev:
	docker-compose up app -d --force-recreate
	$(MAKE) exec-bash