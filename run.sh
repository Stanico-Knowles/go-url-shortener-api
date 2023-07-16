#!/bin/bash

# Load environment variables from .env file
set -a
[ -f .env ] && source .env
set +a

# Run Go project
go run server.go
