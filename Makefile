startDB:
	docker start postgrestulb2
startRedis:
	docker start rdbtulb2

postgres:
	docker run --name postgrestulb4 -p 5432:5432 -e POSTGRES_PASSWORD=tulb -e  POSTGRES_USER=tulb -d postgres

createdb:
	docker exec -it postgrestulb4 createdb --username=tulb --owner=tulb pomodoro

dropdb:
	docker exec -it postgrestulb4 dropdb --username=tulb pomodoro

execdb:
	docker exec -it postgrestulb4 psql -U tulb pomodoro

migratecreate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateupdate:
	migrate create -ext sql -dir db/migration -seq update_schema

migrateup:
	migrate -path db/migration -database "postgresql://tulb:secret@localhost:5432/pomodoro?sslmode=disable" -verbose up

migratedown1:
	migrate -path db/migration -database "postgresql://tulb:secret@localhost:5432/pomodoro?sslmode=disable" -verbose down 1

migrateup1:
	migrate -path db/migration -database "postgresql://tulb:secret@localhost:5432/pomodoro?sslmode=disable" -verbose up 1

sqlc:
	sqlc generate

server:
	go run main.go
.PHONY: postgres createdb dropdb migratecreate migrateup migratedown1 migrateup1 sqlc server startDB startRedis

network:
	docker network create -d bridge mynetwork

redis:
	docker run -d \
	--name rdbtulb2 \
	--hostname redis \
	-p 6379:6379 \
	--network mynetwork \
	redis

exec_redis:
	docker exec -it rdbtulb2 redis-cli
