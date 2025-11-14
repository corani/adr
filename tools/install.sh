#!/bin/bash
dir=$(readlink -f "$(dirname $0)")
export GOBIN="${dir}/bin"
cd "${dir}"

echo "# Found tools:"
tools=$(cat ${dir}/tools.go | grep "_" | cut -d'"' -f2)
for tool in ${tools}; do
    echo "- ${tool}"
done

echo "[CMD] go get ..."
go get ${tools}

echo "[CMD] go mod tidy"
go mod tidy

echo "[CMD] go install ..."
go install ${tools}
