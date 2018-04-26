# apiomat-go

[![Build Status](https://travis-ci.org/philippgille/apiomat-go.svg?branch=master)](https://travis-ci.org/philippgille/apiomat-go) [![Build status](https://ci.appveyor.com/api/projects/status/s8rxuaww5jrmfe21?svg=true)](https://ci.appveyor.com/project/philippgille/apiomat-go) [![Go Report Card](https://goreportcard.com/badge/github.com/philippgille/apiomat-go)](https://goreportcard.com/report/github.com/philippgille/apiomat-go) [![GitHub Releases](https://img.shields.io/github/release/philippgille/apiomat-go.svg)](https://github.com/philippgille/apiomat-go/releases)

Go (golang) packages and CLI for ApiOmat

Under construction!

<img src="https://octodex.github.com/images/constructocat2.jpg" alt="under-construction" width="150"/> [![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

## Contents

- [aoms - Package with shared functionality for sending requests to ApiOmat](#aoms)
- [aomc - Package for handling "customer" resources during design-time](#aomc)
- [aomu - Package for handling "user" resources during runtime](#aomu)
- [aom - CLI for handling both "customer" and "user" resources](#aom)

## aoms

Package aoms contains common functionality for sending HTTP requests to an ApiOmat server.

The "s" in aoms is for "shared", because the package is used by the aomc and aomu packages.
Instead of using this package directly, you should use the packages aomc or aomu instead.

[![GoDoc](https://godoc.org/github.com/philippgille/apiomat-go/aoms?status.svg)](https://godoc.org/github.com/philippgille/apiomat-go/aoms)

## aomc

Package for handling "customer" resources during design-time

### Usage

[![GoDoc](https://godoc.org/github.com/philippgille/apiomat-go/aomc?status.svg)](https://godoc.org/github.com/philippgille/apiomat-go/aomc)

#### Example

```go
package main

import (
    "fmt"

    "github.com/philippgille/apiomat-go/aomc"
)

func main() {
    client := aomc.NewDefaultClient("http://localhost:8080/yambas/rest", "john", "secret", "")
    classes, err := client.GetClasses("TestModule", "")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v", classes) // [{AllowedRolesCreate:[] AllowedRolesGrant:[] ...} {...}]
}
```

### Build

For updating the ApiOmat customer resource types such as `Class`, you can turn example JSON into structs with [gojson](https://github.com/ChimeraCoder/gojson) like this:

```bash
gojson -input "json/class.json" -name "Class" -o "class_struct.go" -pkg "aomc"
```

## aomu

Package for handling "user" resources during runtime

## aom

CLI for handling both "customer" and "user" resources

### Installation

`go get "github.com/philippgille/apiomat-go/aom"`

### Usage

Parameters:

```
  -baseUrl string
        base URL, for example "https://epdemo.apiomat.enterprises/yambas/rest"
  -module string
        Name of the module, for example "Basics" (default "Basics")
  -password string
        password (default "secret")
  -username string
        username (default "apinaut")
```

#### Example:

```bash
$ aom -baseUrl "http://localhost:8080/yambas/rest" -username "john" -password "secret" -module "TestModule"

Version: {"server":"localhost:8080","version":"2.6.2-107E"}
Classes of module TestModule: [{AllowedRolesCreate:[] AllowedRolesGrant:[] ...} {...}]
```
