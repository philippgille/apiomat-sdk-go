package main

import (
	"flag"
	"fmt"

	"github.com/philippgille/apiomat-go/aomc"
)

func main() {
	// CLI arguments
	baseUrl := flag.String("baseUrl", "http://localhost:8080/yambas/rest", `base URL, for example "http://localhost:8080/yambas/rest"`)
	username := flag.String("username", "apinaut", "username")
	password := flag.String("password", "secret", "password")
	module := flag.String("module", "Basics", `Name of the module, for example "Basics"`)
	flag.Parse()

	// Version
	client := aomc.NewDefaultClient(*baseUrl, *username, *password, "")
	version, err := client.GetVersion()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nVersion: %v", version) // {"server":"null:443","version":"2.6.2-107E"}

	// Classes
	classes, err := client.GetClasses(*module, "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nClasses of module %v: %+v", *module, classes) // [{AllowedRolesCreate:[] AllowedRolesGrant:[] ...} {...}]
}
