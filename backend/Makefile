.PHONY:


build:
	go mod download && go build -o ./server ./cmd/server/main.go

docker:
	docker-compose up

migrations-up:
	docker-compose up migrations

migrations-down:
	MIGRATIONS_STATUS=down docker-compose -f docker-compose-down-migrations.yaml up

run: build
	./server

test:
	go test -v ./...

swag:
	docker pull quay.io/goswagger/swagger
	docker run --rm -it --env GOPATH=/go -v $(PWD):/go/src -w /go/src --entrypoint /bin/sh quay.io/goswagger/swagger ./docs/swag.sh

lint:
	golangci-lint run

fmt:
	go fmt ./...