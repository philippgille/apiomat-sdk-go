// Package aomm contains common functionality for sending HTTP requests to an ApiOmat instance.
// Instead of using this package directly, you should use the packages aomc or aomu instead.
package aomm

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AomClient is a client for ApiOmat
// You should create an AomClient object with aomm.NewAomClient()
type AomClient struct {
	// ApiOmat base URL, for example: "https://epdemo.apiomat.enterprises/yambas/rest"
	BaseUrl  string
	username string
	password string
	// Default system is "LIVE"
	system System

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

// NewAomClient creates a new ApiOmat client
func NewAomClient(baseUrl string, username string, password string, system System) AomClient {
	// Remove trailing "/"
	baseUrl = strings.TrimRight(baseUrl, "/")

	return AomClient{
		BaseUrl:    baseUrl,
		username:   username,
		password:   password,
		system:     system,
		httpClient: &http.Client{},
	}
}

// Get sends an HTTP GET request to a URL that consists of the AomClient's base URL and the given path and URL parameters
func (client AomClient) Get(path string, params url.Values) (string, error) {
	// Create URL
	path = "/" + strings.TrimLeft(path, "/")
	url, err := url.Parse(client.BaseUrl + path)
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
		encodedAuth := base64.StdEncoding.EncodeToString([]byte(client.username + ":" + client.password))
		req.Header.Add("Authorization", "Basic "+encodedAuth)
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

// GetVersion returns the version of the ApiOmat instance the AomClient is configured for.
// Example return value: "Yambas REST interface v2.6.2-107E on null:80"
func (client AomClient) GetVersion() (string, error) {
	return client.Get("", nil)
}
