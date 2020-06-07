#!/bin/sh
set -ex

BINARY_PATH=/tmp/networkplan

rm -rf "${BINARY_PATH}"
go build -o "${BINARY_PATH}"
"${BINARY_PATH}" "$@"
rm -rf "${BINARY_PATH}"
