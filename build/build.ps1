$ErrorActionPreference = "Stop"

Write-Output "Environment:"
go version
go env

# Builds
Write-Output "Building"
go build github.com/philippgille/apiomat-sdk-go/aomx
go build github.com/philippgille/apiomat-sdk-go/aomc

# Tests
Write-Output "Running tests"
go test -v github.com/philippgille/apiomat-sdk-go/aomx
go test -v github.com/philippgille/apiomat-sdk-go/aomc
