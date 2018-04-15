package main

import (
	"flag"
	"fmt"

	"github.com/philippgille/apiomat-go/aomc"
	"github.com/philippgille/apiomat-go/aomm"
)

func main() {
	// CLI arguments
	baseUrl := flag.String("baseUrl", "", `base URL, for example "https://epdemo.apiomat.enterprises/yambas/rest"`)
	username := flag.String("username", "apinaut", "username")
	password := flag.String("password", "secret", "password")
	module := flag.String("module", "Basics", "Name of the module")
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
	fmt.Println(classes) // [{5ac5bbd76d79587667be0b40 https://epdemo.apiomat.enterprises/yambas/rest/modules/ArzModule/metamodels/5ac5bbd76d79587667be0b40 StandingOrder} ... ]
}
