$ErrorActionPreference = "Stop"

Write-Output "Environment:"
go version
go env

# Builds
Write-Output "Building"
go build github.com/philippgille/apiomat-go/aoms
go build github.com/philippgille/apiomat-go/aomc
go build github.com/philippgille/apiomat-go/aom

# Tests
Write-Output "Running tests"
go test -v github.com/philippgille/apiomat-go/aoms
go test -v github.com/philippgille/apiomat-go/aomc
