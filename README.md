# api-template

## Getting started

### Prerequisites

Flowgo requires [Go](https://go.dev/) version [1.24](https://go.dev/doc/devel/release#go1.24.0) or above.

### Getting api-template
```shell
go get -u github.com/admgo/api-template
```

### Running api-template
#### 1. write your api to api/openapi/api.yaml file
#### 2. generate handler file
```shell
cd scripts
./api-gen.sh
```
#### 2. run in root dir
```shell
go run cmd/server/main.go
```
