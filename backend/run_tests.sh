#!/bin/bash

# Navigate to the directory containing your Go tests
# Replace this with the path to your Go project's root directory
cd /path/to/your/go/project

# Run the Go tests with verbose output
go test -v ./api

# If your tests are in a specific package, you can specify the package path
# go test -v ./path/to/your/package
