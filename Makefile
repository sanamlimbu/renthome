PACKAGE=renthome

# Names and versions
DOCKER_CONTAINER=$(PACKAGE)-db

# Paths
SERVER = $(CURDIR)/server
BIN = $(CURDIR)/bin
# DB settings
LOCAL_DEV_DB_USER=$(PACKAGE)
LOCAL_DEV_DB_PASS=devdev
LOCAL_DEV_DB_HOST=localhost
LOCAL_DEV_DB_PORT=5438
LOCAL_DEV_DB_DATABASE=${PACKAGE}
DB_CONNECTION_STRING="postgres://$(LOCAL_DEV_DB_USER):$(LOCAL_DEV_DB_PASS)@$(LOCAL_DEV_DB_HOST):$(LOCAL_DEV_DB_PORT)/$(LOCAL_DEV_DB_DATABASE)?sslmode=disable"

# Commands
.PHONY: docker-create
docker-create:
	docker run -d  -p ${LOCAL_DEV_DB_PORT}:5432 --name ${DOCKER_CONTAINER} -e POSTGRES_USER=${LOCAL_DEV_DB_USER} -e POSTGRES_PASSWORD=${LOCAL_DEV_DB_PASS} -e POSTGRES_DB=${PACKAGE} postgres:15

.PHONY: docker-start
docker-start:
	docker start ${DOCKER_CONTAINER}

.PHONY: docker-stop
docker-stop:
	docker stop ${DOCKER_CONTAINER}

.PHONY: docker-remove
docker-remove:
	docker rm ${DOCKER_CONTAINER}

.PHONY: docker-setup
docker-setup:
	docker exec -it $(DOCKER_CONTAINER) psql -U $(PACKAGE) -c 'CREATE EXTENSION IF NOT EXISTS pg_trgm; CREATE EXTENSION IF NOT EXISTS pgcrypto; CREATE EXTENSION IF NOT EXISTS "uuid-ossp";'

.PHONY: db-drop
db-drop:
	bin/migrate -database $(DB_CONNECTION_STRING) -path $(SERVER)/migrations drop -f

.PHONY: db-migrate
db-migrate:
	bin/migrate -database $(DB_CONNECTION_STRING) -path $(SERVER)/migrations up

.PHONY: db-prepare
db-prepare: db-drop db-migrate

.PHONY: db-seed
db-seed:
	cd server && go run seed/main.go db

.PHONY: db-reset
db-reset: db-drop db-migrate db-seed

.PHONY: go-mod-tidy
go-mod-tidy:
	cd server && go mod tidy

.PHONY: go-mod-download
go-mod-download:
	cd server && go mod download

.PHONY: tools-darwin
tools-darwin: go-mod-tidy
	@mkdir -p $(BIN)
	cd server && go generate -tags tools ./tools/...

.PHONY: serve
serve:
	${BIN}/air -c ./server/.air.toml

.PHONY: web-watch
web-watch:
	cd web/ && npm run start

.PHONY: generate
generate:
	$(BIN)/sqlboiler $(BIN)/sqlboiler-psql --wipe --tag boiler --config $(SERVER)/sqlboiler.toml --output $(SERVER)/boiler