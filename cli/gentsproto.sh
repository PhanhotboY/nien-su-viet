#!/bin/bash

# In a bash script, set -e is a command that enables the "exit immediately" option. When this option is set, the script will terminate immediately if any command within the script exits with a non-zero status (indicating an error).
set -e

readonly service="$1"
readonly outPath="./apps/gateway/grpc/genproto/$service-service"

mkdir -p "$outPath"

protoc \
    --plugin=./node_modules/.bin/protoc-gen-ts_proto \
    --proto_path="api/protobuf/$service-service" \
    --ts_proto_opt=nestJs=true \
    --ts_proto_out="$outPath" \
    api/protobuf/$service-service/*.proto
