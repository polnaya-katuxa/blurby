SERVICE_NAME = coursework
PG_MIGRATION = add_events
CH_MIGRATION = add_aggregated_view
KAFKA_BROKERS=broker:29092
PG_MIGRATIONS_DIR = "deployments/migrations/postgres"
CH_MIGRATIONS_DIR = "deployments/migrations/clickhouse"
PG_DSN = postgres://$(SERVICE_NAME):password@127.0.0.1:5436/$(SERVICE_NAME)?sslmode=disable
CH_DSN = tcp://localhost:9000/$(SERVICE_NAME)?username=$(SERVICE_NAME)&password=password&dial_timeout=500ms

install-goose:
	go get github.com/pressly/goose/cmd/goose
	go install github.com/pressly/goose/cmd/goose

install-goimports:
	go install golang.org/x/tools/cmd/goimports@latest

goimports:
	goimports -w .

create-pg-migration:
	goose --dir=$(PG_MIGRATIONS_DIR) create $(PG_MIGRATION) sql

create-ch-migration:
	goose --dir=$(CH_MIGRATIONS_DIR) create $(CH_MIGRATION) go

migrate-pg:
	goose --dir=$(PG_MIGRATIONS_DIR) postgres "$(PG_DSN)" up

migrate-ch:
	go build -o goose-custom $(CH_MIGRATIONS_DIR)/*.go
	./goose-custom -dir $(CH_MIGRATIONS_DIR) -b $(KAFKA_BROKERS) -dsn "$(CH_DSN)" -cmd up

migrate-down-pg:
	goose --dir=$(PG_MIGRATIONS_DIR) postgres "$(PG_DSN)" down

migrate-down-ch:
	go build -o goose-custom $(CH_MIGRATIONS_DIR)/*.go
	./goose-custom -dir $(CH_MIGRATIONS_DIR) -b $(KAFKA_BROKERS) -dsn "$(CH_DSN)" -cmd down

run-docker:
	docker-compose --project-directory deployments up -d

shut-down-docker:
	docker-compose --project-directory deployments stop

_gen-server:
	docker run --rm \
      -v ${PWD}:/local openapitools/openapi-generator-cli generate \
      -i /local/api/swagger.yaml \
      -g go-server \
      -o /local/internal/server/ \
      --additional-properties=outputAsLibrary=true,sourceFolder=.

gen-server: _gen-server goimports

_gen-client:
	docker run --rm \
      -v ${PWD}:/local openapitools/openapi-generator-cli generate \
      -i /local/api/swagger.yaml \
      -g go \
      --git-host=git.iu7.bmstu.ru --git-user-id iu7-kostritsky --git-repo-id iu7-db-course-project-2023-karpovaekaterina-backend/pkg/client \
      -o /local/pkg/client/ \
      --additional-properties=isGoSubmodule=true,sourceFolder=.
	rm pkg/client/go.mod pkg/client/go.sum
	rm -rf pkg/client/test

gen-client: _gen-client goimports

run:
	go run cmd/$(SERVICE_NAME)/main.go -cfg ./configs/local_config.yml

test:
	go test ./... -cover

install-golint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2

research:
	go run cmd/research/main.go

lint:
	golangci-lint run -E dupl,errname,goconst,gofumpt,goimports,wrapcheck \
	--skip-dirs internal/server,internal/mocks,internal/containers --timeout 5m0s

fmt:
	gofumpt -l -w .