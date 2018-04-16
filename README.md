# apiomat-go

Go (golang) packages and CLI for ApiOmat

[![Build Status](https://travis-ci.org/philippgille/apiomat-go.svg?branch=master)](https://travis-ci.org/philippgille/apiomat-go) [![GitHub Releases](https://img.shields.io/github/release/philippgille/apiomat-go.svg)](https://github.com/philippgille/apiomat-go/releases)

Under construction!

<img src="https://octodex.github.com/images/constructocat2.jpg" alt="under-construction" width="150"/> [![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

## Contents

- [aomc - Package for handling "customer" resources during design-time](#aomc)
- [aomu - Package for handling "user" resources during runtime](#aomu)
- [aom - CLI for handling both "customer" and "user" resources](#aom)

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
    aomcClient := aomc.NewAomClient("http://localhost:8080/yambas/rest", "john", "secret", "")
    classes, err := aomcClient.GetClasses("TestModule", "")
    if err != nil {
        // handle error
    }
    fmt.Println(classes) // [{5ac5bbd76d79587667be0b40 http://localhost:8080/yambas/rest/modules/TestModule/metamodels/5ac5bbd76d79587667be0b40 BankUser} ... ]
}
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

{"server":"null:443","version":"2.6.2-107E"}
[{5ac5bbd76d79587667be0b40 http://localhost:8080/yambas/rest/modules/TestModule/metamodels/5ac5bbd76d79587667be0b40 TestClass} {5ac776326d79587667bf8987 http://localhost:8080/yambas/rest/modules/TestModule/metamodels/5ac776326d79587667bf8987 TestClass2}]
```