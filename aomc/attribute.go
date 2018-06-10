package aomc

import (
	"net/url"
	"time"

	"github.com/philippgille/apiomat-go/aomc/dto"
)

// Attribute represents an ApiOmat class attribute, also called "MetaModelAttribute".
// It's a convencience type that contains proper Go types (like `time.Time` instead of
// a Unix epoch timestamp in milliseconds).
//
// Some attributes are exported, but changing their value won't lead
// to a changed value on the server when updating the entity.
// For example, "Created" is set by the server when the attribute is
// created and can never be changed. It's an exported field instead of
// a getter to avoid inconsistency.
type Attribute struct {
	// Read only fields
	ID           string
	URL          url.URL
	Created      time.Time
	LastModified time.Time
	// Writable when creating, not when updating
	Name string
	// Readable and writable
	IsBinary      bool // Corresponds to the "image" field in the JSON
	IsCollection  bool
	IsDeprecated  bool
	IsEmbedded    bool
	IsIgnored     bool
	IsIndexed     bool
	IsMandatory   bool
	IsReadOnly    bool
	IsReference   bool
	IsSensitive   bool
	RoleForUpdate string
	Type          string
	// Not exported
	isFromOtherModule bool
	metaModelURL      url.URL
	refID             string
	// Unused
	//CapitalizedName string
	//LowerCaseName string
}

// GetAttributes returns the attributes of the given class in the given module.
func (client Client) GetAttributes(module string, classID string) ([]Attribute, error) {
	result := []Attribute{}
	rawAttributes, err := client.GetRawAttributes(module, classID)
	if err != nil {
		return nil, err
	}
	for _, rawAttribute := range rawAttributes {
		// Convert existing fields
		elem := ConvertAttributeFromDto(rawAttribute)

		result = append(result, elem)
	}
	return result, nil
}

// GetRawAttributes returns the raw attributes of the given class in the given module.
// "Raw" means the struct is mapped 1:1 to a "MetaModelAttribute" JSON
func (client Client) GetRawAttributes(module string, classID string) ([]dto.Attribute, error) {
	jsonString, err := client.Get("modules/"+module+"/metamodels/"+classID+"/attributes", nil)
	if err != nil {
		return nil, err
	}
	result, err := ConvertRawAttributesFromJSON(jsonString)
	if err != nil {
		return nil, err
	}
	return result, nil
}
