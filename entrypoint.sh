#!/bin/sh -l

COMMAND=$1
VERSION=$2
CONFIG=$3

echo "$CONFIG" | base64 -d > release-automator.yaml

release-automator --version="$VERSION" "$COMMAND"
