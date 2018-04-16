package main

import (
	"flag"
	"fmt"

	"github.com/philippgille/apiomat-go/aomc"
	"github.com/philippgille/apiomat-go/aomm"
)

func main() {
	// CLI arguments
	baseUrl := flag.String("baseUrl", "http://localhost:8080/yambas/rest", `base URL, for example "http://localhost:8080/yambas/rest"`)
	username := flag.String("username", "apinaut", "username")
	password := flag.String("password", "secret", "password")
	module := flag.String("module", "Basics", `Name of the module, for example "Basics"`)
	flag.Parse()

	// Version
	aommClient := aomm.NewAomClient(*baseUrl, "", "", "")
	version, err := aommClient.GetVersion()
	if err != nil {
		panic(err)
	}
	fmt.Println(version) // {"server":"null:443","version":"2.6.2-107E"}

	// Classes
	aomcClient := aomc.NewAomClient(aommClient.BaseUrl, *username, *password, "")
	classes, err := aomcClient.GetClasses(*module, "")
	if err != nil {
		panic(err)
	}
	fmt.Println(classes) // [{5ac5bbd76d79587667be0b40 http://localhost:8080/yambas/rest/modules/TestModule/metamodels/5ac5bbd76d79587667be0b40 TestClass} ... ]
}
