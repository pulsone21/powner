dev:
  go run ./cmd/main.go

build:
  go build -o ./bin/ ./cmd/main.go

run: build
  ./bin/main

db:
  sqlite3 ./db/test.db

