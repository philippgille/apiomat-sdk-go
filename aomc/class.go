package aomc

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/philippgille/apiomat-go/aomc/dto"
)

// Class represents an ApiOmat class, also called "MetaModel".
// It's a convencience type that includes attribute objects in
// contrast to the plain MetaModel JSON from ApiOmat, as well
// as proper Go types (like `time.Time` instead of a timestamp).
//
// Some attributes are exposed, but changing their value won't lead
// to a changed value on the server when updating the entity.
// For example, "CreatedAt" is set by the server when the class is
// created and can never be changed. It's an exposed field instead
// of a getter to avoid inconsistency.
type Class struct {
	// Read only fields
	ID           string
	URL          url.URL
	Created      time.Time
	LastModified time.Time
	// Writable when creating, not when updating
	Name string
	// Readable and writable
	AllowedRolesCreate     []string
	AllowedRolesGrant      []string
	AllowedRolesRead       []string
	AllowedRolesWrite      []string
	Attributes             []Attribute
	IsDeprecated           bool
	IsGlobal               bool
	IsInvisible            bool
	IsRestrictBinaryAccess bool // Corresponds to "restrictResourceAccess" in the JSON
	IsTransient            bool
	RequiredUserRoleCreate string
	RequiredUserRoleGrant  string
	RequiredUserRoleRead   string
	RequiredUserRoleWrite  string
	UseOwnAuth             string
	// Non-exposed
	attributesURL url.URL
	methodsURL    url.URL
	moduleURL     url.URL
	// Unused
	//ExtendsGeoModel bool
}

// Attribute represents an ApiOmat class attribute, also called "MetaModelAttribute".
// It's a convencience type that contains proper Go types (like `time.Time` instead of a timestamp).
//
// Some attributes are exposed, but changing their value won't lead
// to a changed value on the server when updating the entity.
// For example, "CreatedAt" is set by the server when the attribute is
// created and can never be changed. It's an exposed field instead of
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
	// Not exposed
	isFromOtherModule bool
	metaModelURL      url.URL
	refID             string
	// Unused
	//CapitalizedName string
	//LowerCaseName string
}

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
func (client Client) GetRawAttributes(module string, classID string) ([]dto.Attribute, error) {
	var result []dto.Attribute
	jsonString, err := client.Get("modules/"+module+"/metamodels/"+classID+"/attributes", nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't unmarshal the response body")
	}
	return result, nil
}

// GetClasses returns the classes of the given module including their attributes.
func (client Client) GetClasses(module string) ([]Class, error) {
	result := []Class{}
	rawClasses, err := client.GetRawClasses(module)
	if err != nil {
		return nil, err
	}
	for _, rawClass := range rawClasses {
		// Convert existing fields
		elem := ConvertClassFromDto(rawClass)
		// Fetch and assign attributes
		elem.Attributes, err = client.GetAttributes(module, elem.ID)

		result = append(result, elem)
	}
	return result, nil
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
