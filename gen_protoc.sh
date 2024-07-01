# bin/bash

echo "==============Generating Proto==========="
protoc --go_out=src/proto \
    --go-grpc_out=src/proto \
    src/proto/todo.proto

protoc --go_out=src/proto \
    --go-grpc_out=src/proto \
    src/proto/stream.proto

echo "============Done generating Proto========"