package aomc_test

import (
	"fmt"

	"github.com/philippgille/apiomat-go/aomc"
)

var baseURL = "https://apiomat.yourcompany.com/yambas/rest"
var username = "john"
var password = "secret"

func Example() {
	// 1. Create a new Client instance
	client := aomc.NewDefaultClient(baseURL, username, password, "")
	// 2. Call one of the client's methods. For example, to fetch all classes of a module
	classes, err := client.GetClasses("TestModule")
	// 3. Handle the possible error (it's a github.com/pkg/errors error that contains a stack trace.)
	if err != nil {
		panic(err)
	}
	// 4. Work with the result
	fmt.Printf("%+v", classes) // [{AllowedRolesCreate:[] AllowedRolesGrant:[] ...} {...}]
}
