.PHONY: all
all: cmd/generate/generate cmd/serve/serve

cmd/generate/generate: pkg/puzzle/*.go
	go build -o $@ cmd/generate/main.go

cmd/serve/serve: pkg/puzzle/*.go internal/game/*.go
	go build -o $@ cmd/serve/main.go

.PHONY: clean
clean:
	rm cmd/generate/generate
	rm cmd/serve/serve

