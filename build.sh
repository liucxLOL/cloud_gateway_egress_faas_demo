#! /bin/bash
set -ex

CURDIR=$(cd $(dirname $0); pwd)

if [ ! -d "$CURDIR/bin" ]; then
    mkdir -p "$CURDIR/bin"
fi

GOPROXY=https://goproxy.cn,direct GOOS=linux go build -o "$CURDIR/bin/demo"
