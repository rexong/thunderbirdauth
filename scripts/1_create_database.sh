#!/bin/bash

source "$(dirname $0)/0_load_env.sh"

echo "--- Setting up Database---"
DB_PATH="$DATA_DIR/$DB_NAME"
mkdir -p "$DATA_DIR"
if [ $? -ne 0 ]; then
  echo "Error: Failed to create directory $DATA_DIR"
  exit 1
fi

sqlite3 "$DB_PATH" ".quit"
if [ $? -ne 0 ]; then
  echo "Error: Failed to create the SQLite database file."
  exit 1
fi

echo "--- Database Set Up at $DB_PATH ---"

