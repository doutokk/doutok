#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=bff
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}