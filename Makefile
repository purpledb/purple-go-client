GO = go

build:
	$(GO) build -v -mod vendor ./...

fmt:
	gofmt -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

tidy:
	go mod tidy

imports:
	goimports -d $(shell find . -type f -name '*.go' -not -path "./vendor/*")

spruce: tidy fmt imports

test:
	$(GO) test -p 1 -v ./...

run-example-grpc-client:
	go run examples/grpc-client/main.go

run-example-http-client:
	go run examples/http-client/main.go
