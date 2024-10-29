#!/bin/sh

set -ex

cd "$(dirname "$0")"

go build -o "$1/mongobench" ./mongobench.go
