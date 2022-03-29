include .env

## for windows CLI: mingw32-make start/stop/createdb
start: ;
	docker-compose up -d

stop: ;
	docker-compose stop

createdb: ;
	docker exec -it $(CONTAINER_DB_NAME) createdb --username=$(APP_DB_USERNAME) $(APP_DB_NAME)

dropdb: ;
	docker exec -it $(CONTAINER_DB_NAME) dropdb --username=$(APP_DB_USERNAME) $(APP_DB_NAME)

## see https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md

migration = migrate -path $(PATH_DB_MIGRATION) -database "postgresql://$(APP_DB_USERNAME):$(APP_DB_PASSWORD)@$(APP_DB_HOST):$(APP_DB_PORT)/$(APP_DB_NAME)?sslmode=disable"

## v: migration's name
## with arg: mingw32-make createmigration name=test_migration
createmigration: ;
	@[ "${name}" ] || ( echo ">> Name of migration is not set"; exit 1 )
	migrate create -ext sql -dir $(PATH_DB_MIGRATION) -seq $(name)

## v: migrate to version
## with arg: mingw32-make migrateup v=5
migrateup: ;
	$(migration) up $(v)

## v: migrate to version
## with arg: mingw32-make migratedown v=2
migratedown:
	$(migration) down $(v)

.PHONY: start stop createdb dropdb migrateup migratedown createmigration