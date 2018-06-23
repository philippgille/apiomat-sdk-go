package aomc

import (
	"net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/philippgille/apiomat-sdk-go/aomc/dto"
)

// Class represents an ApiOmat class, also called "MetaModel".
// It's a convencience type that includes attribute objects in contrast to
// the plain MetaModel JSON from ApiOmat, as well as proper Go types
// (like `time.Time` instead of a Unix epoch timestamp in milliseconds).
//
// Some attributes are exported, but changing their value won't lead
// to a changed value on the server when updating the entity.
// For example, "Created" is set by the server when the class is
// created and can never be changed. It's an exported field instead
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
	AllowedRolesCreate         []string
	AllowedRolesGrant          []string
	AllowedRolesRead           []string
	AllowedRolesWrite          []string
	Attributes                 []Attribute
	AuthImplStatus             AuthImplStatus
	IsCopyRolesToBinAttributes bool // Corresponds to "restrictResourceAccess" in the JSON
	IsDeprecated               bool
	IsGlobal                   bool
	IsInvisible                bool
	IsTransient                bool
	RequiredRoleCreate         UserRole
	RequiredRoleGrant          UserRole
	RequiredRoleRead           UserRole
	RequiredRoleWrite          UserRole
	// Non-exported
	attributesURL url.URL
	methodsURL    url.URL
	moduleURL     url.URL
	// Unused
	//ExtendsGeoModel bool
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

// GetClassByName returns the class in the given module with the given name including its attributes.
func (client Client) GetClassByName(module string, name string) (Class, error) {
	rawClass, err := client.GetRawClassByName(module, name)
	if err != nil {
		return Class{}, err
	}
	// Convert existing fields
	result := ConvertClassFromDto(rawClass)
	// Fetch and assign attributes
	result.Attributes, err = client.GetAttributes(module, result.ID)
	return result, nil
}

// GetRawClasses returns the raw classes of the given module.
// "Raw" means the struct is mapped 1:1 to a "MetaModel" JSON, without embedded attribute structs for example.
func (client Client) GetRawClasses(module string) ([]dto.Class, error) {
	jsonString, err := client.Get("modules/"+module+"/metamodels", nil)
	if err != nil {
		return nil, err
	}
	result, err := ConvertRawClassesFromJSON(jsonString)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRawClassByName returns the raw class of the given module with the given name.
// "Raw" means the struct is mapped 1:1 to a "MetaModel" JSON, without embedded attribute structs for example.
func (client Client) GetRawClassByName(module string, name string) (dto.Class, error) {
	// ApiOmat doesn't have an endpoint for fetching a single Class by name
	// so we get all classes and return the one with the given name.
	rawClasses, err := client.GetRawClasses(module)
	if err != nil {
		return dto.Class{}, err
	}
	for _, rawClass := range rawClasses {
		if name == rawClass.Name {
			return rawClass, nil
		}
	}
	return dto.Class{}, errors.New("No class was found in module " + module + " with the name " + name)
}
