createdb:
	docker exec -it mirket createdb --username=root --owner=root -e user
	
dropdb:
	docker exec -it mirket dropdb user

migrate_up:
	migrate -path migrations -database "postgresql://root:password@localhost:5432/user?sslmode=disable" -verbose up

migrate_down:
	migrate -path migrations -database "postgresql://root:password@localhost:5432/user?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

mock:
	mockgen -destination=internal/user/mock/mock.go -package=mock github.com/wralith/mirket/src/userservice/internal/user/repo Repo

.PHONY:
	createdb dropdb migrate_up migrate_down sqlc proto server