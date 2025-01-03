#!/bin/bash

set -e

ignore_file="$PWD/.coverignore"
coverage_file="$PWD/coverage.out"

excluded_packages=$(grep -v '^$' "$ignore_file" | tr '\n' '|' | sed 's/|$//')
packages=$(go list ./... | grep -vE "$excluded_packages" | tr '\n' ' ' | sed 's/ $//')

go test -cover -covermode=atomic -coverprofile="$coverage_file" $packages
