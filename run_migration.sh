#!/bin/bash

# Load environment variables from .env file
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
else
  echo ".env file not found!"
  exit 1
fi

# Check if CONN is set
if [ -z "$CONN" ]; then
  echo "CONN is not set in the .env file!"
  exit 1
fi

# Check if an argument is provided (up or down)
if [ -z "$1" ]; then
  echo "Usage: $0 [up|down]"
  exit 1
fi

# Determine the migration direction and folder
DIRECTION=$1
MIGRATIONS_FOLDER="./sql/schema"

# Run goose migrations based on the direction
if [ "$DIRECTION" = "up" ]; then
  goose -dir $MIGRATIONS_FOLDER postgres "$CONN" up
elif [ "$DIRECTION" = "down" ]; then
  goose -dir $MIGRATIONS_FOLDER postgres "$CONN" down
else
  echo "Invalid argument: $DIRECTION. Use 'up' or 'down'."
  exit 1
fi

# Check if the migration command was successful
if [ $? -eq 0 ]; then
  echo "Migrations applied successfully!"
else
  echo "Failed to apply migrations."
  exit 1
fi
