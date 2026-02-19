.PHONY: build test lint run clean

build:
	cd cmd/vecforge-server && go build -o vecforge .

test:
	cd cmd/vecforge-server && go test -v .

lint:
	golangci-lint run ./...

run: build
	./cmd/vecforge-server/vecforge

clean:
	rm -f cmd/vecforge-server/vecforge
