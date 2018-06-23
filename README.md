# apiomat-sdk-go

[![Build Status](https://travis-ci.org/philippgille/apiomat-sdk-go.svg?branch=master)](https://travis-ci.org/philippgille/apiomat-sdk-go) [![Build status](https://ci.appveyor.com/api/projects/status/61qhp3dm6a7ukrtq/branch/master?svg=true)](https://ci.appveyor.com/project/philippgille/apiomat-sdk-go/branch/master) [![Go Report Card](https://goreportcard.com/badge/github.com/philippgille/apiomat-sdk-go)](https://goreportcard.com/report/github.com/philippgille/apiomat-sdk-go) [![GitHub Releases](https://img.shields.io/github/release/philippgille/apiomat-sdk-go.svg)](https://github.com/philippgille/apiomat-sdk-go/releases)

Go (golang) SDK for ApiOmat

Under construction!

<img src="https://octodex.github.com/images/constructocat2.jpg" alt="under-construction" width="150"/> [![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

## Contents

- [aomx - Package with shared functionality for sending requests to ApiOmat](#aomx)
- [aomc - Package for handling "customer" resources during design-time](#aomc)
- [aomu - Package for handling "user" resources during runtime](#aomu)

## aomx

Package aomx contains common functionality for sending HTTP requests to an ApiOmat server.

The "x" in aomx is for "cross" (cutting functionality), because the package is used by the aomc and aomu packages.
Instead of using this package directly, you should use the packages `aomc` or `aomu` instead.

[![GoDoc](https://godoc.org/github.com/philippgille/apiomat-sdk-go/aomx?status.svg)](https://godoc.org/github.com/philippgille/apiomat-sdk-go/aomx)

## aomc

Package for handling "customer" resources during design-time

### Usage

[![GoDoc](https://godoc.org/github.com/philippgille/apiomat-sdk-go/aomc?status.svg)](https://godoc.org/github.com/philippgille/apiomat-sdk-go/aomc)

#### Example

```go
package main

import (
    "fmt"

    "github.com/philippgille/apiomat-sdk-go/aomc"
)

func main() {
    client := aomc.NewDefaultClient("https://apiomat.yourcompany.com/yambas/rest", "john", "secret", "")
    classes, err := client.GetClasses("TestModule")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v", classes) // [{AllowedRolesCreate:[] AllowedRolesGrant:[] ...} {...}]
}
```

### Build

For updating the ApiOmat customer resource types such as `dto.Class`, you can turn example JSON into structs with [gojson](https://github.com/ChimeraCoder/gojson) like this:

```bash
gojson -input "dto/class.json" -name "Class" -o "dto/class.go" -pkg "dto"
```

## aomu

Package for handling "user" resources during runtime
