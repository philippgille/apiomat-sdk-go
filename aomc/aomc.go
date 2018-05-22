/*
Package aomc is for handling ApiOmat "customer" resources during design-time.

For example, you can fetch all classes of an ApiOmat module.

You can use the package like this:
1. Create a new Client instance with aomc.NewDefaultClient()
2. Call for example the client's GetClasses() method to fetch all classes of a given module
*/
package aomc

import (
	"github.com/philippgille/apiomat-go/aoms"
)

// Client is a client for ApiOmat customer resources
type Client struct {
	// Embedded anonymous type
	aoms.Client
}

// NewDefaultClient creates a new ApiOmat client with an underlying aoms.DefaultClient.
// username, password and system may be empty.
// If username or password are empty, no HTTP Authorization header is set in the HTTP request.
// If system is empty, no "X-Apiomat-System" header is set in the HTTP request, leading to "LIVE" being used as default by ApiOmat.
func NewDefaultClient(baseUrl string, username string, password string, system aoms.System) Client {
	return Client{
		Client: aoms.NewDefaultClient(baseUrl, username, password, system),
	}
}

// NewClient creates a new ApiOmat client with a given client that implements the aoms.Client interface.
func NewClient(client aoms.Client) Client {
	return Client{
		Client: client,
	}
}

// Methods are grouped by resource (for example "class") and reside in their own files (for example "class.go")
