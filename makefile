live/templ:
	templ generate --watch --proxy="http://localhost:5551" --cmd="go run ./cmd/main.go" -v -v

live/tailwind:
	npx tailwindcss -i ./public/static/styles/input.css -o ./public/static/styles/base.css --minify --watch

live/server:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "go build -o tmp/bin/main" --build.bin "tmp/bin/main" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

live: 
	make -j2 live/templ live/tailwind
