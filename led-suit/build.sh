#!/usr/bin/env bash

if [ -z "$1" ]
then 
    # build all
    for op in GR RS BK GB
    do
        ./"$0" $op
    done
else
    if [ "$1" != "GR" ] && [ "$1" != "RS" ] && [ "$1" != "BK" ] && [ "$1" != "GB" ]
    then
        echo "unknown option $1"
        exit 1
    fi

    echo "building $1"

    PACKAGE="ledsuit"
    LDFLAGS=(
        "-X '${PACKAGE}/config.Selection=$1'"
    )

    tinygo build -ldflags="${LDFLAGS[*]}" -target pico -o "$1".uf2 ./cmd/pico
fi
