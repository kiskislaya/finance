run:
	go run ./cmd/finance/main.go
build:
	go build -o ./bin/finance ./cmd/finance/main.go
get-deps:
	go get github.com/jackc/pgx/v5/pgxpool
	go get github.com/redis/go-redis/v9
install-deps:
	go install github.com/pressly/goose/v3/cmd/goose@latest

create-migration:
	goose -dir=migrations create $(name) sql
migrate-up:
	goose -dir=migrations postgres "user=finance password=finance host=localhost port=5433 sslmode=disable" up
migrate-down:
	goose -dir=migrations postgres "user=finance password=finance host=localhost port=5433 sslmode=disable" down