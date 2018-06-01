/*
Package aomc is for handling ApiOmat "customer" resources during design-time.

For example, you can fetch all classes of an ApiOmat module.

You can use the package like this:
1. Create a new Client instance with aomc.NewDefaultClient()
2. Call for example the client's GetClasses() method to fetch all classes of a given module
*/
package aomc

import (
	"github.com/philippgille/apiomat-go/aomx"
)

// Client is a client for ApiOmat customer resources
type Client struct {
	// Embedded anonymous type
	aomx.Client
}

// NewDefaultClient creates a new ApiOmat client with an underlying aomx.DefaultClient.
// username, password and system may be empty.
// If username or password are empty, no HTTP Authorization header is set in the HTTP request.
// If system is empty, no "X-Apiomat-System" header is set in the HTTP request, leading to "LIVE" being used as default by ApiOmat.
func NewDefaultClient(baseURL string, username string, password string, system aomx.System) Client {
	return Client{
		Client: aomx.NewDefaultClient(baseURL, username, password, system),
	}
}

// NewClient creates a new ApiOmat client with a given client that implements the aomx.Client interface.
func NewClient(client aomx.Client) Client {
	return Client{
		Client: client,
	}
}

// Methods are grouped by resource (for example "class") and reside in their own files (for example "class.go")
