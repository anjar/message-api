#!/bin/bash

set -e

# Copy .env file
if [ ! -f "/app/simpleapi/.env" ]; then
	cp /app/simpleapi/.env.example /app/simpleapi/.env
fi

if [ ! -f "/app/simpleapi/go.mod" ]; then
	echo "[build.sh: Initialize go.mod for $APP_NAME]"
	go mod init simpleapi
fi

# Run compiled service
go run /app/simpleapi/main.go "$@"