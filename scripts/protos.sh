#!/usr/bin/env bash

function main() {
    local proto
    proto="$1"; shift

    protoc \
        --go_out=./internal \
        --go_opt=paths=source_relative \
        --go-grpc_out=./internal \
        --go-grpc_opt=paths=source_relative "${proto}"
}

main api/registry/registry.proto
main api/sandhog/sandhog.proto
