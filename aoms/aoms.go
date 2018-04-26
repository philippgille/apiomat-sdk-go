/*
Package aoms contains common functionality for sending HTTP requests to an ApiOmat server.

The "s" in aoms is for "shared", because the package is used by the aomc and aomu packages.
Instead of using this package directly, you should use the packages aomc or aomu instead.

If you need some functionality that's not implemented by the aomc and aomu packages you can use this package like this:
1. Create a client with aoms.NewDefaultClient()
2. Call the client's Get() method to send an HTTP GET request to an ApiOmat URL of your choice
3. The returned string is JSON which you can parse and work with
*/
package aoms

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client is the interface for ApiOmat clients.
// Use the DefaultClient, which is an implementation of this interface, or implement your own.
// You can create your own type and embed DefaultClient to extend its functionality.
type Client interface {
	GetVersion() (string, error)
	Get(path string, params url.Values) (string, error)
}

// DefaultClient is a client for ApiOmat, which implements the Client interface.
// You should create a DefaultClient object with aoms.NewDefaultClient().
type DefaultClient struct {
	// ApiOmat base URL, for example: "https://epdemo.apiomat.enterprises/yambas/rest"
	baseUrl  string
	username string
	password string
	// This client doesn't set a default value for system,
	// but ApiOmat servers treat requests without system as requests to the "LIVE" system
	system System

	// Safe for concurrent use,
	// so we can be sure that when using DefaultClient as value receiver mutation of any field doesn't lead to deadlocks
	httpClient *http.Client
}

// System is the ApiOmat system to be used.
type System string

func (s *System) String() string {
	return string(*s)
}

// ApiOmat system values
const (
	Live    System = "LIVE"
	Staging System = "STAGING"
	Test    System = "TEST"
)

// NewDefaultClient creates a new ApiOmat client in the form of a DefaultClient.
// baseUrl must be in the form of: "https://epdemo.apiomat.enterprises/yambas/rest"
// username, password and system may be empty.
// If username or password are empty, no HTTP Authorization header is set in the HTTP request.
// If system is empty, no "X-Apiomat-System" header is set in the HTTP request, leading to "LIVE" being used as default by ApiOmat.
func NewDefaultClient(baseUrl string, username string, password string, system System) DefaultClient {
	// Remove trailing "/"
	baseUrl = strings.TrimRight(baseUrl, "/")

	return DefaultClient{
		baseUrl:    baseUrl,
		username:   username,
		password:   password,
		system:     system,
		httpClient: &http.Client{},
	}
}

// Get sends an HTTP GET request to a URL that consists of the DefaultClient's base URL and the given path and URL parameters
// path may be be empty, params may be nil.
// Part of the Client interface implementation.
func (client DefaultClient) Get(path string, params url.Values) (string, error) {
	// Create URL
	path = "/" + strings.TrimLeft(path, "/")
	url, err := url.Parse(client.baseUrl + path)
	if err != nil {
		return "", err
	}
	url.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return "", err
	}

	// Set headers
	req.Header.Add("Accept", "application/json")
	if client.system != "" {
		req.Header.Add("X-Apiomat-System", client.system.String())
	}
	if client.username != "" && client.password != "" {
		req.SetBasicAuth(client.username, client.password)
	}

	// Send request
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read and return body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// GetVersion returns the version of the ApiOmat instance the DefaultClient is configured for.
// Example return value: "Yambas REST interface v2.6.2-107E on null:80"
// Part of the Client interface implementation.
func (client DefaultClient) GetVersion() (string, error) {
	return client.Get("", nil)
}
