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

// NewAomClient creates a new ApiOmat client
func NewAomClient(baseUrl string, username string, password string, system aomm.System) AomClient {
	return AomClient{
		client: aomm.NewAomClient(baseUrl, username, password, system),
	}
}

// GetClasses returns
func (client AomClient) GetClasses(module string, system aomm.System) ([]Class, error) {
	jsonString, err := client.client.Get("modules/"+module+"/metamodels", nil)
	if err != nil {
		return nil, err
	}
	var classes []Class
	json.Unmarshal([]byte(jsonString), &classes)
	return classes, nil
}
