JSDIR=./web/static/src/scripts
CSSDIR=./web/static/src/styles

.PHONY: all
all: cmd/generate/generate cmd/serve/serve web/static/dist/main.js web/static/dist/main.css

.PHONY: run
run: all
	cd cmd/serve && ./serve

cmd/generate/generate: pkg/puzzle/*.go
	go build -o $@ cmd/generate/main.go

cmd/serve/serve: pkg/puzzle/*.go internal/game/*.go
	go build -o $@ cmd/serve/main.go

web/static/dist/main.js: $(JSDIR)/*.js $(JSDIR)/**/*.js
	cat $(JSDIR)/*.js $(JSDIR)/**/*.js > web/static/dist/main.js
	tr -d [:space:] < web/static/dist/main.js > web/static/dist/main.min.js

web/static/dist/main.css: $(CSSDIR)/*.css
	cat $(CSSDIR)/*.css > web/static/dist/main.css

.PHONY: clean
clean:
	rm cmd/generate/generate
	rm cmd/serve/serve
	rm web/static/dist/main.js
	rm web/static/dist/main.css
