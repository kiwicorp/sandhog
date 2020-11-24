#!/usr/bin/env bash

function main() {
    local proto
    proto="$1"; shift

    protoc \
        --proto_path=./api \
        --go_out=./internal/api \
        --go_opt=paths=source_relative \
        --go-grpc_out=./internal/api \
        --go-grpc_opt=paths=source_relative "${proto}"
}

main selftechio/naas/common.proto
main selftechio/naas/naas.proto
