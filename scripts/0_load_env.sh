#!/bin/bash

ENV_FILENAME=".env"

echo "--- Loading Env File ---"
if [ -f "$ENV_FILENAME" ]; then
  source "$ENV_FILENAME"
else
  echo "Error: $ENV_PATH file not found. Exiting."
  exit 1
fi
if [ -z "$DATA_DIR" ] || [ -z "$DB_NAME" ]; then
  echo "Error: DB_DIR or DB_NAME variables are missing or empty in $ENV_FILENAME."
  echo "Please ensure $ENV_FILENAME defines: DB_DIR and DB_NAME"
  exit 1
fi
echo "--- Loaded ENV variables from $ENV_FILENAME ---"
