// Package aomc is for handling ApiOmat "customer" resources during design-time.
// For example, you can fetch all classes of an ApiOmat module.
package aomc

import (
	"encoding/json"

	"github.com/philippgille/apiomat-go/aomm"
)

// AomClient is a client for ApiOmat
type AomClient struct {
	client aomm.AomClient
}

// Class is the representation of an ApiOmat class (sometimes called MetaModel)
type Class struct {
	Id   string `json:"id"`
	Href string `json:"href"`
	Name string `json:"name"`
}

// NewAomClient creates a new ApiOmat client.
// username, password and system may be empty.
// If username or password are empty, no HTTP Authorization header is set in the HTTP request.
// If system is empty, no "X-Apiomat-System" header is set in the HTTP request, leading to "LIVE" being used as default by ApiOmat.
func NewAomClient(baseUrl string, username string, password string, system aomm.System) AomClient {
	return AomClient{
		client: aomm.NewAomClient(baseUrl, username, password, system),
	}
}

// GetClasses returns the classes of the given module and system.
// system may be empty.
// If system is empty, no "X-Apiomat-System" header is set in the HTTP request, leading to "LIVE" being used as default by ApiOmat.
// Example return value: [{5ac5bbd76d79587667be0b40 https://epdemo.apiomat.enterprises/yambas/rest/modules/TestModule/metamodels/5ac5bbd76d79587667be0b40 BankUser} {5ac776326d79587667bf8987 https://epdemo.apiomat.enterprises/yambas/rest/modules/TestModule/metamodels/5ac776326d79587667bf8987 StandingOrder}]
func (client AomClient) GetClasses(module string, system aomm.System) ([]Class, error) {
	jsonString, err := client.client.Get("modules/"+module+"/metamodels", nil)
	if err != nil {
		return nil, err
	}
	var classes []Class
	json.Unmarshal([]byte(jsonString), &classes)
	return classes, nil
}
