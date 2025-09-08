#!/bin/bash
set -e

go tool oapi-codegen -generate "gin-server"  -o "../internal/handler/openapi_api.gen.go" -package "handler"  "../api/openapi/api.yaml"
go tool oapi-codegen -generate "types"  -o "../internal/handler/types.gen.go" -package "handler"  "../api/openapi/api.yaml"
go tool oapi-codegen -generate "spec"  -o "../internal/handler/openapi_spec.gen.go" -package "handler"  "../api/openapi/api.yaml"