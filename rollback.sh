#!/bin/bash

# Read the stable version from the file
stable_version=$(cat stable_version.txt)

# Checkout the stable version
git checkout $stable_version

# Build and deploy the application
go build main.go
./main &