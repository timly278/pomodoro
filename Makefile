postgres:
	docker run --name postgrestulb2 -p 5432:5432 -e POSTGRES_PASSWORD=tulb -e  POSTGRES_USER=root -d postgres

createdb:
	docker exec -it postgrestulb2 createdb --username=root --owner=root pomodoro

dropdb:
	docker exec -it postgrestulb2 createdb dropdb pomodoro

execdb:
	docker exec -it postgrestulb2 psql -U root pomodoro

migratecreate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateupdate:
	migrate create -ext sql -dir db/migration -seq update_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:tulb@localhost:5432/pomodoro?sslmode=disable" -verbose up

migratedown1:
	migrate -path db/migration -database "postgresql://root:tulb@localhost:5432/pomodoro?sslmode=disable" -verbose down 1

migrateup1:
	migrate -path db/migration -database "postgresql://root:tulb@localhost:5432/pomodoro?sslmode=disable" -verbose up 1

sqlc:
	sqlc generate

server:
	go run ./cmd/main.go
.PHONY: postgres createdb dropdb migratecreate migrateup migratedown1 migrateup1 sqlc server
