#!/bin/bash

set -e

# Copy .env file
if [ ! -f "/app/warungpintar/.env" ]; then
	cp /app/warungpintar/.env.example /app/warungpintar/.env
fi

if [ ! -f "/app/warungpintar/go.mod" ]; then
	echo "[build.sh: Initialize go.mod for $APP_NAME]"
	go mod init warungpintar
fi

# Run compiled service
go run /app/warungpintar/main.go "$@"