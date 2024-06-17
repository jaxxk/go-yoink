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

# start server
nohup ./build.sh &

# Wait for the server to start
echo "Waiting for the server to start..."
until curl -s http://localhost:8080/healthz; do
  sleep 1
done

echo "Server is up and running."

# start test
go test .

