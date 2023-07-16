#!/bin/bash

# Check if .env file exists
if [ -f .env ]; then
  # Load environment variables from .env file
  set -a
  source .env
  set +a

  # Run Go project
  go run server.go
else
  echo "Error: .env file not found."
fi
