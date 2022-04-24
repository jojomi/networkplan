#!/bin/sh
set -ex

BINARY_PATH=/tmp/networkplan

rm -rf "${BINARY_PATH}"
go build -o "${BINARY_PATH}"


# build README for Github
io --input "{\"binary_path\": \"${BINARY_PATH}\"}" --allow-exec --template docu/README.tpl.md --output README.md


"${BINARY_PATH}" "$@"

rm -rf "${BINARY_PATH}"
