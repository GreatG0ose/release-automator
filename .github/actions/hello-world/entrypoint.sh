#!/bin/sh -l

COMMAND=$1
CONFIG=$3
VERSION=$3


echo "$CONFIG" | base64 -d > release-automator.yaml

go run ./cmd/release-automator --version=$VERSION $COMMAND

echo cat .build/tweet.txt >> $GITHUB_OUTPUT