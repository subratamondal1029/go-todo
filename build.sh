#!/bin/bash

set -euo pipefail

echo "Running database migrations..."
dbmate up

echo "Building application..."
mkdir -p bin
GOOS=linux GOARCH=amd64 go build -o bin/todo cmd/todo/main.go

echo "Application built successfully."