
#!/bin/bash

# In a bash script, set -e is a command that enables the "exit immediately" option. When this option is set, the script will terminate immediately if any command within the script exits with a non-zero status (indicating an error).
set -e

readonly modulePath="github.com/phanhotboy/nien-su-viet/libs/pkg"
readonly outPath="./libs/pkg"

mkdir -p "./libs/pkg/grpc/genproto"

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
    api/proto/events/*.proto
