DB_PORT:=3999
LATEST := v1

server/on:
	DB_PORT=$(DB_PORT) go run cmd/main.go

database/on:
	export DB_PORT=$(DB_PORT)&& \
	docker compose -f ./tools/database.yml up -d &&\
	sleep 20 &&\
	mysqldef -uroot -proot -h localhost -P 3999 db1 < migration/$(LATEST).sql

database/off:
	export DB_PORT=$(DB_PORT)&& \
	docker compose -f ./tools/database.yml down --remove-orphans \
