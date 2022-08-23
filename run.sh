#! /bin/bash
set -ex

CURDIR=$(cd $(dirname $0); pwd)

export _FAAS_RUNTIME_PORT="8000"

if [ ! -d "$CURDIR/bin" ]; then
    mkdir -p "$CURDIR/bin"
fi

$CURDIR/build.sh
exec "$CURDIR/bin/demo"
