migrate-file:
	migrate create -ext sql -dir migrations ${filename}

swagger:
	swag init --parseDependency --parseInternal

install:
	sh install.sh

run:
	go run main.go start --config=./config/local/config.yaml

wire:
	wire ./internal

compose:
	docker compose -f docker-compose.local.yaml up -d

compose-dev:
	docker compose -f docker-compose.dev.yaml up -d

mock:
	mockgen --build_flags=--mod=mod --destination=./internal/domain/interfaces/user/mocks/mock.go base_service/internal/domain/interfaces/user UserCommandRepository,UserQueryRepository,CacheRepository

test:
	go test -v ./...

