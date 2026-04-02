#!/bin/bash

# In a bash script, set -e is a command that enables the "exit immediately" option. When this option is set, the script will terminate immediately if any command within the script exits with a non-zero status (indicating an error).
set -e

readonly service="$1"
readonly modulePath="github.com/phanhotboy/nien-su-viet/apps/$service"
readonly outPath="./apps/${service}"

mkdir -p "./apps/$service/internal/shared/grpc/genproto"

# https://stackoverflow.com/questions/13616033/install-protocol-buffers-on-windows
# https://dev.to/techschoolguru/how-to-define-a-protobuf-message-and-generate-go-code-4g4e
protoc \
  --proto_path="api/proto" \
  --go_out="$outPath" \
  --go_opt=module="$modulePath" \
  --go-grpc_out="$outPath" \
  --go-grpc_opt=module="$modulePath" \
  --go-grpc_opt=require_unimplemented_servers=false \
    api/proto/common/*.proto \
    api/proto/${service}_service/*.proto
