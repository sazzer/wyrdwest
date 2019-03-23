.DEFAULT_GOAL=all
GO111MODULE=on

clean:
	rm -rf output

all: get-deps lint test
	go build -o output/wyrdwest-service github.com/sazzer/wyrdwest/service/cmd/service

run: get-deps lint test
	go run github.com/sazzer/wyrdwest/service/cmd/service

lint: get-deps
	golangci-lint run --enable-all

test: get-deps
	go test ./...

get-deps:
	go get
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
