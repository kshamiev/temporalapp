# Инициализация
export DIR := $(realpath .)
export SERVICE_NAME := sample
export IMAGE_NAME := "tn/$(SERVICE_NAME):latest"

## Creating .env file from template, if file not exists
#ifeq ("$(wildcard $(DIR)/.env)","")
#	RSP1 := $(shell cp -v $(DIR)/.env_example $(DIR)/.env)
#endif
#include $(DIR)/.env

# Set OS config
ifeq ($(OS),Windows_NT)
	ETC := win
else
	ETC := unix
endif

default: help

# Зависимости
dep:
	go mod tidy
	go mod vendor
.PHONY: dep

# Сваггер
swag:
	swag i --pd --parseVendor --packageName swagger -o internal/swagger -ot go
	swag i --pd --parseVendor -o doc -ot json
.PHONY: swag

# Генерация моков ко внешним сервисам
mock:
	mockgen --build_flags=--mod=mod -package=mocks -destination=test/mocks/sample_client.go application/services/sample/pb SampleClient
.PHONY: mock

# FMT & GOIMPORT
fmt:
	goimports -local application -w internal
.PHONY: fmt

# Linters
lint:
	# golangci-lint cache clean
	golangci-lint run -c .golangci.yml
.PHONY: lint

test-u:
	go test ./... -tags unit
.PHONY: test-u

test-i:
	go test ./... -tags integration -v
.PHONY: test-i

test-b:
	go test ./... -bench=. -benchtime=10x -count 3
.PHONY: test-b

test: test-u test-i test-b
.PHONY: test

# Запуск в режиме разработки
run:
	go build -o bin/app .
	bin/app -c config.yml
.PHONY: run

# Запуск в режиме отладки
dev: dep swag fmt lint test-u
	go build -o bin/app .;
	bin/app -c config.yml;
.PHONY: dev

# Статус миграции
mig-st:
	pgmigrate -c config_boiler.yml -o version
.PHONY: mig-st

# Миграция на одну позицию назад
mig-dn:
	pgmigrate -c config_boiler.yml -o down
.PHONY: mig-dn

# Миграция вперёд до конца
mig-up:
	pgmigrate -c config_boiler.yml -o up
.PHONY: mig-up

# Миграция назад до конца (сброс всех миграций)
mig-reset:
	pgmigrate -c config_boiler.yml -o reset
.PHONY: mig-reset

# Инженеринг типов и методов по существующей структуре БД
entity:
	sqlboiler -c etc/config_boiler.yml -p $@ -o internal/repo/$@ --add-soft-deletes --tag db --no-tests --wipe psql
	genboiler -c etc/config_boiler.yml -entity internal/repo/$@ # генерация UpsertAll, nested sets, native load для типов БД и прочее безобразие с тегами
	goimports -w internal/repo/$@/.
.PHONY: entity

## Встраивание внешней статики в приложение.
## Программой находящейся внутри пакета генерируются исходные коды в которые
## встраиваются статические файлы с данными.
## EMBED_STATIC_RESOURCES - Группы ресурсов и папками самих ресурсов.
##     Формат: group_name1:path/to/folder1,group_name2:/absolut/path/to/folder2
## Относительные пути берутся от текущей рабочей директории
embed:
	@echo "Встроенные ресурсы"
	@export EMBED_STATIC_RESOURCES=www:www; \
	go run cmd/embedder/main.go
.PHONY: embed

# Создание образа
dki:
	docker build --no-cache -t $(IMAGE_NAME) .
.PHONY: dki

#--mount type=bind,src=$(DIR)/www,dst=/home/app/www \
#--mount type=bind,src=$(DIR)/log,dst=/home/app/log \
# Запуск контейнера
dkr:
	docker stop $(SERVICE_NAME) || true && docker rm $(SERVICE_NAME) || true
	docker run --rm --name $(SERVICE_NAME) \
		-p 8080:8080 \
		-p 7070:7070 \
		-v $(DIR)/www:/home/app/www \
		-v $(DIR)/log:/home/app/log \
		-v $(DIR)/etc:/home/app/etc \
		-v $(DIR)/data:/home/app/data \
		-d $(IMAGE_NAME)
.PHONY: dkr

# Запуск окружения
dk-start:
	@echo "Конфигурация метрик"
	@if [ ! -d "etc/grafana" ]; then \
	cp -r doc/${ETC}/etc . || true; \
	fi
	# Postgres
	docker stop $(SERVICE_NAME)-postgres || true && docker rm $(SERVICE_NAME)-postgres || true
	docker run --rm --name $(SERVICE_NAME)-postgres \
	-p 5435:5432 \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_PASSWORD=postgres \
	-e POSTGRES_DB=$(SERVICE_NAME) \
	-d postgres:15

	# Jaeger
	docker stop $(SERVICE_NAME)-jaeger || true && docker rm $(SERVICE_NAME)-jaeger || true
	docker run --rm --name $(SERVICE_NAME)-jaeger \
	-p 16686:16686 -p 14268:14268 -p 4317:4317 -p 4318:4318 -p 14250:14250 \
	-d jaegertracing/all-in-one:1.56

	# RabbitMQ
	docker stop $(SERVICE_NAME)-rabbitmq || true && docker rm $(SERVICE_NAME)-rabbitmq || true
	docker run --rm --name $(SERVICE_NAME)-rabbitmq \
	-p 5672:5672 -p 15672:15672 \
	-d rabbitmq:management

	# Minio (s3)
	docker stop $(SERVICE_NAME)-minio || true && docker rm $(SERVICE_NAME)-minio || true
	docker run --rm --name $(SERVICE_NAME)-minio --network host \
	-p 9000:9000 -p 9190:9090 \
	-e MINIO_ROOT_USER="admin" -e MINIO_ROOT_PASSWORD="xxx-xxx-xxx" \
	-v $(DIR)/data/minio:/data \
	-d quay.io/minio/minio server /data --console-address ":9190"

	# Redis
	docker stop $(SERVICE_NAME)-redis || true && docker rm $(SERVICE_NAME)-redis || true
	docker run --rm --name $(SERVICE_NAME)-redis \
	-p 6379:6379 \
	-d redis

ifeq (${ETC},win)
	# Prometheus
	docker stop $(SERVICE_NAME)-prometheus || true && docker rm $(SERVICE_NAME)-prometheus || true
	docker run --rm --name $(SERVICE_NAME)-prometheus \
	-v $(DIR)/etc/prometheus/prometheus.yaml:/etc/prometheus/prometheus.yml \
	-p 9090:9090 \
	-d quay.io/prometheus/prometheus:latest

	# Grafana
	docker stop $(SERVICE_NAME)-grafana || true && docker rm $(SERVICE_NAME)-grafana || true
	docker run --rm --name $(SERVICE_NAME)-grafana \
	-e GF_SECURITY_ADMIN_USER="admin" \
	-e GF_SECURITY_ADMIN_PASSWORD="prometheus" \
	-e GF_USERS_ALLOW_SIGN_UP="false" \
	-v $(DIR)/etc/grafana/provisioning:/etc/grafana/provisioning \
	-p 3000:3000 \
	-d grafana/grafana:latest
else ifeq (${ETC},unix)
	# Prometheus
	docker stop $(SERVICE_NAME)-prometheus || true && docker rm $(SERVICE_NAME)-prometheus || true
	docker run --rm --name $(SERVICE_NAME)-prometheus --network host \
	-v $(DIR)/etc/prometheus/prometheus.yaml:/etc/prometheus/prometheus.yml \
	-d quay.io/prometheus/prometheus:latest

	# Grafana
	docker stop $(SERVICE_NAME)-grafana || true && docker rm $(SERVICE_NAME)-grafana || true
	docker run --rm --name $(SERVICE_NAME)-grafana --network host \
	-e GF_SECURITY_ADMIN_USER="admin" \
	-e GF_SECURITY_ADMIN_PASSWORD="prometheus" \
	-e GF_USERS_ALLOW_SIGN_UP="false" \
	-v $(DIR)/etc/grafana/provisioning:/etc/grafana/provisioning \
	-d grafana/grafana:latest
endif

.PHONY: dk-start

dk-stop:
	docker stop $(SERVICE_NAME)-postgres
	docker stop $(SERVICE_NAME)-jaeger
	docker stop $(SERVICE_NAME)-rabbitmq
	docker stop $(SERVICE_NAME)-minio
	docker stop $(SERVICE_NAME)-redis
	docker stop $(SERVICE_NAME)-prometheus
	docker stop $(SERVICE_NAME)-grafana
.PHONY: dk-stop

# Help
h:
	@echo "Usage: make [target]"
	@echo "  target is:"
	@echo "       dep	- Обновление зависимостей"
	@echo "      swag	- Генерация документации swagger"
	@echo "      mock	- Генерация моков ко внешним приложениям"
	@echo "       fmt	- Форматирование кодовой базы"
	@echo "      lint	- Линтеры"
	@echo "    test-u	- Unit тесты"
	@echo "    test-i	- Интеграционные тесты"
	@echo "    test-b	- Тесты производительности"
	@echo "    test		- Запуск всех тестов"
	@echo "       run	- Запуск в режиме разработки"
	@echo "       dev	- Запуск в режиме отладки"
	@echo "    mig-st	- Статус миграции"
	@echo "    mig-dn	- Миграция на одну позицию назад"
	@echo "    mig-up	- Миграция вперёд до конца"
	@echo " mig-reset	- Миграция назад до конца (сброс всех миграций)"
	@echo "    entity	- Инженеринг типов и методов по существующей структуре БД"
	@echo "  dk-start	- Запуск служб в докер контейнерах (окружения)"
	@echo "   dk-stop	- Остановка запущенных служб (окружения)"
	@echo "       dki	- Создание docker образа"
	@echo "       dkr	- Запуск сервиса в докер контейнере"
.PHONY: h
help: h
.PHONY: help
