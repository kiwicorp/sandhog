#!/usr/bin/env bash

function main() {
    local out
    out="$1"; shift

    mkdir -p bin/

    go build \
		-o "${out}" \
		-mod vendor \
		-race
}

main $@
