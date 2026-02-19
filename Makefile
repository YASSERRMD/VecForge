.PHONY: build test lint run docker deploy clean

build:
	cd cmd/vecforge-server && go build -o vecforge .

test:
	cd cmd/vecforge-server && go test -v ./...

lint:
	golangci-lint run ./...

run: build
	./cmd/vecforge-server/vecforge

docker:
	docker build -t vecforge -f docker/Dockerfile .

docker-up:
	docker-compose -f docker/docker-compose.yml up -d

deploy:
	flyctl deploy

clean:
	rm -f cmd/vecforge-server/vecforge
