package aomc

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// GetRawClasses returns the raw classes of the given module.
// "Raw" means the struct is mapped 1:1 to a "MetaModel" JSON, without embedded attribute structs for example.
// Example return value: [{AllowedRolesCreate:[] AllowedRolesGrant:[] ...} {...}]
func (client Client) GetRawClasses(module string) ([]RawClass, error) {
	jsonString, err := client.Get("modules/"+module+"/metamodels", nil)
	if err != nil {
		return nil, err
	}
	var rawClasses []RawClass
	err = json.Unmarshal([]byte(jsonString), &rawClasses)
	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't unmarshal the response body")
	}
	return rawClasses, nil
}

// GetRawAttributes returns the raw attributes of the given class.
// "Raw" means the struct is mapped 1:1 to a "MetaModelAttribute" JSON
func (client Client) GetRawAttributes(class RawClass) ([]RawAttribute, error) {
	panic("not implemented yet")
}
