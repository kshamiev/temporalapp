# temporalapp

sample temporal app

Требования: установлено: `go 1.23.2`, `Makefile`

## Первые шаги
Установите [Temporal](https://temporal.io/setup/install-temporal-cli):
```shell
apt-get install temporal
```

Установите [protoc](https://grpc.io/docs/protoc-installation/):
```shell
apt-get install protobuf
```

Установите [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/):
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Установите [protoc-gen-go-temporal](https://github.com/cludden/protoc-gen-go-temporal):
```shell
go install github.com/cludden/protoc-gen-go-temporal/cmd/protoc-gen-go_temporal@latest
```

Установите линтер [workflowcheck](https://pkg.go.dev/go.temporal.io/sdk/contrib/tools/workflowcheck#section-readme):
```shell
go install go.temporal.io/sdk/contrib/tools/workflowcheck@latest
```
