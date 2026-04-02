#!/bin/bash

# In a bash script, set -e is a command that enables the "exit immediately" option. When this option is set, the script will terminate immediately if any command within the script exits with a non-zero status (indicating an error).
set -e

readonly service="$1"
readonly outPath="./libs/nsv-interfaces/grpc/genproto"

mkdir -p "$outPath"
rm -rf "$outPath/*"

protoc \
    --plugin=./node_modules/.bin/protoc-gen-ts_proto \
    --proto_path="api/proto" \
    --ts_proto_opt=nestJs=true \
    --ts_proto_out="$outPath" \
    api/proto/${service}_service/*.proto
