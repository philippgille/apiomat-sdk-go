package aomc

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// GetClasses returns the classes of the given module.
// Example return value: [{AllowedRolesCreate:[] AllowedRolesGrant:[] ...} {...}]
func (client Client) GetClasses(module string) ([]Class, error) {
	jsonString, err := client.Get("modules/"+module+"/metamodels", nil)
	if err != nil {
		return nil, err
	}
	var classes []Class
	err = json.Unmarshal([]byte(jsonString), &classes)
	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't unmarshal the response body")
	}
	return classes, nil
}
