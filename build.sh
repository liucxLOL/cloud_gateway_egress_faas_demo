#! /bin/bash
set -ex

CURDIR=$(cd $(dirname $0); pwd)

if [ ! -d "$CURDIR/bin" ]; then
    mkdir -p "$CURDIR/bin"
fi

go build -o "$CURDIR/bin/demo"
