/*
Package aomc is for handling ApiOmat "customer" resources during design-time.

For example, you can fetch all classes of an ApiOmat module.

Have a look at the GoDoc code examples to see how to use this package.

All returned errors are github.com/pkg/errors errors and contain a stack trace.

The returned structs differ from the JSON that the ApiOmat server returns,
because some attribute names in the JSON are misleading, verbose or inconsistent.
*/
package aomc

import (
	"github.com/philippgille/apiomat-sdk-go/aomx"
)

// Client is a client for ApiOmat customer resources.
//
// It has an embedded type, which is an implementation of the aomx.Client interface,
// so you can use the interface's methods as well (like `Get(path string, params url.Values) (string, error)`).
type Client struct {
	// Embedded anonymous type
	aomx.Client
}

// NewDefaultClient creates a new ApiOmat client with an underlying aomx.DefaultClient.
//
// Username, password and system may be empty.
// If username or password are empty, no HTTP Authorization header is set in the HTTP request.
// If system is empty, no "X-Apiomat-System" header is set in the HTTP request, leading to "LIVE" being used as default by ApiOmat.
func NewDefaultClient(baseURL string, username string, password string, system aomx.System) Client {
	return Client{
		Client: aomx.NewDefaultClient(baseURL, username, password, system),
	}
}

// NewClient creates a new ApiOmat client with a given client that implements the aomx.Client interface.
//
// This is mainly useful for writing tests or when you need to customize the client's behavior.
func NewClient(client aomx.Client) Client {
	return Client{
		Client: client,
	}
}

// Methods are grouped by resource (for example "class") and reside in their own files (for example "class.go")
