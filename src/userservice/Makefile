createdb:
	docker exec -it mirket createdb --username=root --owner=root -e user
	
dropdb:
	docker exec -it mirket dropdb user

migrate_up:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/user?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/user?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run cmd/userservice/main.go

.PHONY:
	createdb dropdb migrate_up migrate_down sqlc proto server mock