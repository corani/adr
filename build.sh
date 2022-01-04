#!/bin/bash
BUILTBY="$USER"
BUILTAT=$(date --iso-8601=seconds --utc)
VERSION=$(git describe --tags --abbrev=0)
VERSION="${VERSION#"v"}-SNAPSHOT"
COMMIT=$(git rev-parse HEAD)

echo "[CMD] mkdir -p bin"
mkdir -p bin

echo "[CMD] go build ..."
go build -ldflags="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${BUILTAT} -X main.builtBy=${BUILTBY}" -o bin/adr .

echo "[CMD] bin/adr version"
bin/adr version
