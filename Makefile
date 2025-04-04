export PROTOC_GEN_GO_TEMPORAL_VERSION=v1.14.3
export TEMPORAL_DEBUG=true

dep:
	go mod tidy
	go mod vendor
.PHONY: dep

gen-temporal:
	@protoc \
     -I ./proto \
     --go_out=../ \
     --go_opt=paths=import \
     --plugin=protoc-gen-go-temporal \
     --go_temporal_out=../ \
     --go_temporal_opt="cli-categories=true" \
     --go_temporal_opt="cli-enabled=true" \
     --go_temporal_opt="workflow-update-enabled=true" \
     proto/common.proto proto/processing.proto proto/customer.proto proto/checkout.proto
.PHONY: gen-temporal

gen-server:
	@protoc \
     -I ./proto \
     --go_out=../ \
     --go-grpc_out=../ \
     proto/server.proto
.PHONY: gen-server

temporal-dev-server:
	temporal server start-dev \
      --dynamic-config-value "frontend.enableUpdateWorkflowExecution=true" \
      --dynamic-config-value "frontend.enableUpdateWorkflowExecutionAsyncAccepted=true" \
      --ui-port 8080
.PHONY: temporal-dev-server

worker:
	go run cmd/worker/main.go worker
.PHONY: worker

server:
	go run cmd/server/main.go
.PHONY: server

# make -j3 all
all: temporal-dev-server worker server