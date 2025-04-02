# Инициализация
export DIR := $(realpath .)
export TEMPORAL_DEBUG=true

default: help

# Help
h:
	@echo "Usage: make [target]"
	@echo "  target is:"
	@echo "                     dep	- Обновление зависимостей"
	@echo "     temporal-dev-server	- Запуск сервера temporal"
	@echo "                  worker - Регистрация и режим работы workflow"
	@echo "                   start - Команда для запуска нового экземпляра workflow"
.PHONY: h
help: h
.PHONY: help

# Зависимости
dep:
	go mod tidy
	go mod vendor
.PHONY: dep

##########################################3

temporal-dev-server:
	temporal server start-dev \
      --dynamic-config-value "frontend.enableUpdateWorkflowExecution=true" \
      --dynamic-config-value "frontend.enableUpdateWorkflowExecutionAsyncAccepted=true" \
      --ui-port 8080
.PHONY: temporal-dev-server

worker:
	go run worker/main.go
.PHONY: worker

server:
	go run starter/main.go
.PHONY: server

# make -j start
start: temporal-dev-server worker server
.PHONY: start
