.DEFAULT_GOAL=all
GO111MODULE=on

clean:
	rm -rf output

all: get-deps lint test
	go build -o output/wyrdwest-service github.com/sazzer/wyrdwest/service/cmd/service

run: get-deps lint test
	go run github.com/sazzer/wyrdwest/service/cmd/service

lint: get-deps
	golangci-lint run --enable-all --disable lll

test: get-deps lint
	go test -covermode=count -coverprofile=output/coverage.out github.com/sazzer/wyrdwest/service/...

cover-report: test
	go tool cover -html=output/coverage.out

get-deps:
	go get
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	go install github.com/golang/mock/mockgen
	go generate github.com/sazzer/wyrdwest/service/...
	mkdir -p output
