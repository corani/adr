#!/bin/bash
echo "[CMD] mkdir -p bin"
mkdir -p bin

echo "[CMD] go build ..."
go build -o bin/adr .

echo "[CMD] bin/adr version"
bin/adr version
