dev: swag
  go run ./cmd/main.go

swag:
  swag fmt --dir ./internal/server,./internal
  swag init --dir ./internal/server,./internal,./internal/server/response --parseDependency --parseInternal

templ: 
  wgo -file .templ templ generate

css:
    npx tailwindcss -i ./public/static/styles/input.css -o ./public/static/styles/base.css --minify --watch

watch_dev:
  wgo -file .templ templ generate :: npx tailwindcss -i ./public/static/styles/input.css -o ./public/static/styles/base.css --minify :: go run ./cmd/main.go

build:
  go build -o ./bin/ ./cmd/main.go

run: build
  ./bin/main

db:
  sqlite3 ./db/test.db

del_db:
  rm -rf ./db/test.db
  touch ./db/test.db

test:
  go test ./... -v

