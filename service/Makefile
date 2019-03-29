.DEFAULT_GOAL=all
GO111MODULE=on

clean: get-packr
	rm -rf output
	rm -rf vendor
	packr clean

all: get-deps lint test
	go build -o output/wyrdwest-service github.com/sazzer/wyrdwest/service/cmd/service

run: get-deps lint test
	go run github.com/sazzer/wyrdwest/service/cmd/service

lint: get-deps
	golangci-lint run --enable-all --disable lll --skip-dirs integration

test: get-deps lint
	go test -covermode=count -coverprofile=output/coverage.out github.com/sazzer/wyrdwest/service/...

cover-report: test
	go tool cover -html=output/coverage.out

get-deps: get-packr
	-go get
	-go get github.com/golangci/golangci-lint/cmd/golangci-lint
	-go install github.com/golang/mock/mockgen
	go generate github.com/sazzer/wyrdwest/service/...
	mkdir -p output

get-packr:
	-go get github.com/gobuffalo/packr/packr

docker: clean lint test
	go mod vendor
	packr
	docker build -t uk.co.grahamcox.wyrdwest.service/wyrdwest-service:latest .
	rm -rf vendor
	packr clean

integration: docker run-integration

run-integration:
	go mod vendor
	docker-compose -f docker-compose.integration.yml -p wyrdwest_integration down
	docker-compose -f docker-compose.integration.yml -p wyrdwest_integration up --exit-code-from=wyrdwest-integration --build --force-recreate --always-recreate-deps
	rm -rf vendor

local-integration: 
	go test -v github.com/sazzer/wyrdwest/service/integration/...
