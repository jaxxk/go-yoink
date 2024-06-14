#!/bin/bash

# Run Goose Migrations
./run_migration.sh up

# Check if the migration command was successful
if [ $? -eq 0 ]; then
  echo "Migrations applied successfully!"
else
  echo "Failed to apply migrations."
  exit 1
fi

# Run the tests
go test -v .
