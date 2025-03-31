# Инициализация
export DIR := $(realpath .)
export TEMPORAL_DEBUG=true

default: help

# Help
h:
	@echo "Usage: make [target]"
	@echo "  target is:"
	@echo "                     dep	- Обновление зависимостей"
	@echo "                   start - Запустить всё и сразу"
.PHONY: h
help: h
.PHONY: help

# Зависимости
dep:
	go mod tidy
	go mod vendor
.PHONY: dep

##########################################3

gen-temporal:
	protoc \
     -I ./proto \
     --go_out=../ \
     --go_opt=paths=import \
     --plugin=protoc-gen-go-temporal \
     --go_temporal_out=../ \
     --go_temporal_opt="cli-categories=true" \
     --go_temporal_opt="cli-enabled=true" \
     --go_temporal_opt="workflow-update-enabled=true" \
     proto/temporal.proto
.PHONY: gen-temporal

gen-server:
	protoc \
     -I ./proto \
     --go_out=../ \
     --go-grpc_out=../ \
     --plugin=protoc-gen-go-temporal \
     proto/server.proto
.PHONY: gen-server

temporal-dev-server:
	temporal server start-dev \
      --dynamic-config-value "frontend.enableUpdateWorkflowExecution=true" \
      --dynamic-config-value "frontend.enableUpdateWorkflowExecutionAsyncAccepted=true" \
      --db-filename data/temporal.txt \
      --ui-port 8080
.PHONY: temporal-dev-server

worker:
	go run worker/main.go worker
.PHONY: worker

server:
	go run server/main.go
.PHONY: server

# make -j start
start: temporal-dev-server worker server
.PHONY: start

