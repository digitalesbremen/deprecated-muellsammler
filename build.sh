#!/usr/bin/env zsh

declare -r BINARY_NAME="bremen_trash"

# CGO_ENABLED=0   -> Disable interoperate with C libraries -> speed up build time! Enable it, if dependencies use C libraries!
# GOOS=linux      -> compile to linux because scratch docker file is linux
# GOARCH=amd64    -> because, hmm, everthing works fine with 64 bit :)
# -a              -> force rebuilding of packages that are already up-to-date.
# -o gpio-test-x  -> force to build an executable gpio-test-x file (instead of default https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies)

go test ./...

echo "Building amd64 binary '$BINARY_NAME-amd64'..."
env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o $BINARY_NAME-amd64

echo "Building arm binary '$BINARY_NAME-arm'..."
env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -a -o $BINARY_NAME-arm