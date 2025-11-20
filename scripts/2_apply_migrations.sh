#!/bin/bash

source "$(dirname $0)/0_load_env.sh"

DB_PATH="$DATA_DIR/$DB_NAME"
MIGRATIONS_DIR="assets/migrations"

for file in $MIGRATIONS_DIR/*.up.sql; do
  echo "Applying migration: $file"
  sqlite3 "$DB_PATH" < "$file"
  if [ $? -ne 0 ]; then
    echo "Error: Applying $file. Stopping..."
    exit 1
  fi
done

echo "Database migrations complete for $DB_PATH"
