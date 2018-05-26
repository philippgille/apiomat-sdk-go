package aomc

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/philippgille/apiomat-go/aomc/dto"
)

// GetRawClasses returns the raw classes of the given module.
// "Raw" means the struct is mapped 1:1 to a "MetaModel" JSON, without embedded attribute structs for example.
// Example return value: [{AllowedRolesCreate:[] AllowedRolesGrant:[] ...} {...}]
func (client Client) GetRawClasses(module string) ([]dto.Class, error) {
	var result []dto.Class
	jsonString, err := client.Get("modules/"+module+"/metamodels", nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't unmarshal the response body")
	}
	return result, nil
}

// GetRawAttributes returns the raw attributes of the given class.
// "Raw" means the struct is mapped 1:1 to a "MetaModelAttribute" JSON
func (client Client) GetRawAttributes(module string, classId string) ([]dto.Attribute, error) {
	var result []dto.Attribute
	jsonString, err := client.Get("modules/"+module+"/metamodels/"+classId+"/attributes", nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't unmarshal the response body")
	}
	return result, nil
}
