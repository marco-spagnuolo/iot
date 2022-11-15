#!/bin/sh

set -e

RED='\033[0;31m'
NC='\033[0m' # No Color
PATH_TO_MAIN="$1"

export CGO_ENABLED=0

# Build
echo "Running go mod download, to get packages ..."
go mod download

echo "Running go test ..."
go test -coverprofile=test.coverage -v -timeout "${GO_TEST_TIMEOUT:-1m}" ./...
go tool cover -func=test.coverage

if [ -f .cover.toml ]
then
  echo "Checking coverage threshold..."
  go-coverage-threshold
else
  echo "No, .cover.toml found, skipping go-coverage-threshold"
fi

echo "Building app in "${PATH_TO_MAIN:-./main.go}" and putting it at /tmp/app"
go build -o /tmp/app "${PATH_TO_MAIN:-./main.go}"
