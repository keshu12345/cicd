#!/bin/bash

# Read the current version from the file
current_version=$(cat stable_version.txt)

# Checkout the previous stable version
git checkout HEAD~1

# Update the stable version file
echo "v1.0.0" > stable_version.txt

# Build and deploy the application
go build main.go
./main &