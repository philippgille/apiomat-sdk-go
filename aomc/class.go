package aomc

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/philippgille/apiomat-go/aomc/dto"
	"github.com/philippgille/apiomat-go/aoms"
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
	Id           string
	Href         url.URL
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
	attributesHref url.URL
	methodsHref    url.URL
	moduleHref     url.URL
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
	Id           string
	Href         url.URL
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
	metaModelHref     url.URL
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
		elem.Attributes, err = client.GetAttributes(module, elem.Id)

		result = append(result, elem)
	}
	return result, nil
}

// GetAttributes returns the attributes of the given class in the given module.
func (client Client) GetAttributes(module string, classId string) ([]Attribute, error) {
	result := []Attribute{}
	rawAttributes, err := client.GetRawAttributes(module, classId)
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

// ConvertClassFromDto converts a raw class of type dto.Class into the convenience type aomc.Class.
// It does NOT fetch and assign the class attributes, it just converts existing fields.
func ConvertClassFromDto(rawClass dto.Class) Class {
	result := Class{
		// First go through all fields of the raw class that can be copied
		AllowedRolesCreate:     rawClass.AllowedRolesCreate,
		AllowedRolesGrant:      rawClass.AllowedRolesGrant,
		AllowedRolesRead:       rawClass.AllowedRolesRead,
		AllowedRolesWrite:      rawClass.AllowedRolesWrite,
		Id:                     rawClass.ID,
		IsGlobal:               rawClass.IsGlobal,
		IsDeprecated:           rawClass.Deprecated,
		IsInvisible:            rawClass.IsInvisible,
		IsRestrictBinaryAccess: rawClass.RestrictResourceAccess,
		IsTransient:            rawClass.IsTransient,
		Name:                   rawClass.Name,
		RequiredUserRoleCreate: rawClass.RequiredUserRoleCreate,
		RequiredUserRoleGrant:  rawClass.RequiredUserRoleGrant,
		RequiredUserRoleRead:   rawClass.RequiredUserRoleRead,
		RequiredUserRoleWrite:  rawClass.RequiredUserRoleWrite,
		UseOwnAuth:             rawClass.UseOwnAuth,
	}
	// Then those that need special attention
	result.attributesHref = *aoms.MustUrl(url.Parse(rawClass.AttributesHref))
	result.Created = time.Unix(0, rawClass.CreatedAt*int64(time.Millisecond))
	result.Href = *aoms.MustUrl(url.Parse(rawClass.Href))
	result.LastModified = time.Unix(0, rawClass.LastModifiedAt*int64(time.Millisecond))
	result.methodsHref = *aoms.MustUrl(url.Parse(rawClass.MethodsHref))
	result.moduleHref = *aoms.MustUrl(url.Parse(rawClass.ModuleHref))

	return result
}

// ConvertAttributeFromDto converts a raw attribute of type dto.Attribute into the convenience type aomc.Attribute.
func ConvertAttributeFromDto(rawAttribute dto.Attribute) Attribute {
	result := Attribute{
		// First go through all fields of the raw class that can be copied
		IsDeprecated:      rawAttribute.Deprecated,
		Id:                rawAttribute.ID,
		IsBinary:          rawAttribute.Image,
		IsCollection:      rawAttribute.IsCollection,
		IsEmbedded:        rawAttribute.IsEmbeddedObject,
		isFromOtherModule: rawAttribute.AddedFromOtherModule,
		IsIgnored:         rawAttribute.JSONIgnore,
		IsIndexed:         rawAttribute.IsIndexed,
		IsMandatory:       rawAttribute.IsMandatory,
		IsReadOnly:        rawAttribute.ReadOnly,
		IsReference:       rawAttribute.IsReference,
		IsSensitive:       rawAttribute.SensibleData,
		Name:              rawAttribute.Name,
		refID:             rawAttribute.RefID,
		RoleForUpdate:     rawAttribute.RoleForUpdate,
		Type:              rawAttribute.Type,
	}
	// Then those that need special attention
	result.Created = time.Unix(0, rawAttribute.CreatedAt*int64(time.Millisecond))
	result.Href = *aoms.MustUrl(url.Parse(rawAttribute.Href))
	result.LastModified = time.Unix(0, rawAttribute.LastModifiedAt*int64(time.Millisecond))
	result.metaModelHref = *aoms.MustUrl(url.Parse(rawAttribute.MetaModelHref))

	return result
}
