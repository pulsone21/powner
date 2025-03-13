include .env
export

# Run templ generation in watch mode
templ:
	templ generate --watch --proxy="http://localhost:"$(PORT) --open-browser=false -v

# Run air for Go hot reload
server:
	air \
  --build.cmd "go build -o ./bin/main ./cmd/main.go" \
  --build.bin "./bin/main" \
  --build.delay "100" \
  --build.exclude_dir "node_modules" \
  --build.include_ext "go" \
  --build.stop_on_error "false" \
  --misc.clean_on_exit true

# Watch Tailwind CSS changes
tailwind:
	tailwindcss -i ./assets/css/input.css -o ./public/static/styles/base.css --watch

swag:
	swag fmt --dir ./internal/server,./internal
	swag init --dir ./internal/server,./internal,./internal/server/response --parseDependency --parseInternal

dev: 
	make -j4 templ server tailwind swag
