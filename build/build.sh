#!/bin/bash

set -euxo pipefail

go version
go env

# Builds
go build github.com/philippgille/apiomat-sdk-go/aomx
go build github.com/philippgille/apiomat-sdk-go/aomc

# Tests
go test -v github.com/philippgille/apiomat-sdk-go/aomx
go test -v github.com/philippgille/apiomat-sdk-go/aomc
