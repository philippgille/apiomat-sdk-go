# apiomat-go

Go (golang) packages and CLI for ApiOmat

Under construction!

<img src="https://octodex.github.com/images/constructocat2.jpg" alt="under-construction" width="150"/> [![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

## aomc

Package for handling "customer" resources during design-time

### Usage

```go
package main

import (
    "fmt"

    "github.com/philippgille/apiomat-go/aomc"
)

func main() {
    aomcClient := aomc.NewAomClient("https://epdemo.apiomat.enterprises/yambas/rest", "john", "secret", "")
    classes, err := aomcClient.GetClasses("TestModule", "")
    if err != nil {
        // handle error
    }
    fmt.Println(classes) // [{5ac5bbd76d79587667be0b40 https://epdemo.apiomat.enterprises/yambas/rest/modules/TestModule/metamodels/5ac5bbd76d79587667be0b40 BankUser} ... ]
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
        Name of the module (default "Basics")
  -password string
        password (default "secret")
  -username string
        username (default "apinaut")
```

Example:

```bash
$ aom -baseUrl "https://epdemo.apiomat.enterprises/yambas/rest" -username "john" -password "secret" -module "TestModule"

{"server":"null:443","version":"2.6.2-107E"}
[{5ac5bbd76d79587667be0b40 https://epdemo.apiomat.enterprises/yambas/rest/modules/TestModule/metamodels/5ac5bbd76d79587667be0b40 BankUser} {5ac776326d79587667bf8987 https://epdemo.apiomat.enterprises/yambas/rest/modules/TestModule/metamodels/5ac776326d79587667bf8987 StandingOrder}]
```