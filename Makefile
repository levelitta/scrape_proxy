LOCAL_BIN:=$(CURDIR)/bin
GOBIN=$(LOCAL_BIN)

.PHONY: .bin-deps
.bin-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

.PHONY: generate
generate: .bin-deps .generate

.PHONY: generate-fast
generate-fast: .generate

.PHONY: .generate
.generate:
	GOBIN=$(LOCAL_BIN) protoc -I ./proto \
		./proto/*.proto \
		--go_out=pkg \
		--go-grpc_out=pkg \
		--plugin protoc-gen-go="$(GOBIN)/protoc-gen-go"

	go generate ./...

.PHONY: test
test:
	go test -v ./...
	cd ./pkg/api && go test -v ./...

.PHONY: test-race
test-race:
	go test --race -v ./...
	cd ./pkg/api && go test --race -v ./...


.PHONY: run
run:
	SCRAPE_PROXY_ADDR=:9010 go run ./cmd/scrape_proxy/main.go