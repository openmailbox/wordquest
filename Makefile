.PHONY: all
all: cmd/generate/generate cmd/serve/serve web/static/dist/main.js

cmd/generate/generate: pkg/puzzle/*.go
	go build -o $@ cmd/generate/main.go

cmd/serve/serve: pkg/puzzle/*.go internal/game/*.go
	go build -o $@ cmd/serve/main.go

web/static/dist/main.js: web/static/src/**/*.js
	npm run build

.PHONY: clean
clean:
	rm cmd/generate/generate
	rm cmd/serve/serve
	rm web/static/dist/main.js

