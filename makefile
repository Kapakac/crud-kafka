include .env

migration = migrate -path db/migration -database "postgresql://$(APP_DB_USERNAME):$(APP_DB_PASSWORD)@localhost:5464/$(APP_DB_NAME)?sslmode=disable" -verbose

start: ;
	docker-compose up -d

stop: ;
	docker-compose stop

createdb: ;
	docker exec -it $(CONTAINER_DB_NAME) createdb --username=$(APP_DB_USERNAME) $(APP_DB_NAME)

dropdb: ;
	docker exec -it $(CONTAINER_DB_NAME) dropdb --username=$(APP_DB_USERNAME) $(APP_DB_NAME)

migrateup: ;
	$(migration) up

migratedown:
	$(migration) down

.PHONY: start stop createdb dropdb migrateup migratedown