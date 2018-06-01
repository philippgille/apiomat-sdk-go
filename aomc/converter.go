package aomc

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/philippgille/apiomat-go/aomc/dto"
	"github.com/philippgille/apiomat-go/aoms"
	"github.com/pkg/errors"
)

// ConvertRawClassesFromJSON converts a JSON array into a slice of dto.Class objects.
// Having an extra function that wraps json.Unmarshal(...) allows us to call the function in tests
// and change the JSON library later without having to change the tests.
func ConvertRawClassesFromJSON(jsonString string) ([]dto.Class, error) {
	var result []dto.Class
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't unmarshal the JSON: %v", jsonString)
	}
	return result, nil
}

// ConvertRawAttributesFromJSON converts a JSON array into a slice of dto.Attribute objects.
// Having an extra function that wraps json.Unmarshal(...) allows us to call the function in tests
// and change the JSON library later without having to change the tests.
func ConvertRawAttributesFromJSON(jsonString string) ([]dto.Attribute, error) {
	var result []dto.Attribute
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't unmarshal the JSON: %v", jsonString)
	}
	return result, nil
}

// ConvertClassFromDto converts a raw class of type dto.Class into the convenience type aomc.Class.
// It does NOT fetch and assign the class attributes, it just converts existing fields.
func ConvertClassFromDto(rawClass dto.Class) Class {
	return Class{
		// Exported
		AllowedRolesCreate: rawClass.AllowedRolesCreate,
		AllowedRolesGrant:  rawClass.AllowedRolesGrant,
		AllowedRolesRead:   rawClass.AllowedRolesRead,
		AllowedRolesWrite:  rawClass.AllowedRolesWrite,
		AuthImplStatus:     AuthImplStatus(rawClass.UseOwnAuth),
		Created:            time.Unix(0, rawClass.CreatedAt*int64(time.Millisecond)),
		ID:                 rawClass.ID,
		IsCopyRolesToBinAttributes: rawClass.RestrictResourceAccess,
		IsGlobal:                   rawClass.IsGlobal,
		IsDeprecated:               rawClass.Deprecated,
		IsInvisible:                rawClass.IsInvisible,
		IsTransient:                rawClass.IsTransient,
		LastModified:               time.Unix(0, rawClass.LastModifiedAt*int64(time.Millisecond)),
		Name:                       rawClass.Name,
		RequiredRoleCreate:         UserRole(rawClass.RequiredUserRoleCreate),
		RequiredRoleGrant:          UserRole(rawClass.RequiredUserRoleGrant),
		RequiredRoleRead:           UserRole(rawClass.RequiredUserRoleRead),
		RequiredRoleWrite:          UserRole(rawClass.RequiredUserRoleWrite),
		URL:                        *aoms.MustUrl(url.Parse(rawClass.Href)),
		// Unexported
		attributesURL: *aoms.MustUrl(url.Parse(rawClass.AttributesHref)),
		methodsURL:    *aoms.MustUrl(url.Parse(rawClass.MethodsHref)),
		moduleURL:     *aoms.MustUrl(url.Parse(rawClass.ModuleHref)),
	}
}

// ConvertAttributeFromDto converts a raw attribute of type dto.Attribute into the convenience type aomc.Attribute.
func ConvertAttributeFromDto(rawAttribute dto.Attribute) Attribute {
	return Attribute{
		// Exported
		Created:       time.Unix(0, rawAttribute.CreatedAt*int64(time.Millisecond)),
		ID:            rawAttribute.ID,
		IsDeprecated:  rawAttribute.Deprecated,
		IsBinary:      rawAttribute.Image,
		IsCollection:  rawAttribute.IsCollection,
		IsEmbedded:    rawAttribute.IsEmbeddedObject,
		IsIgnored:     rawAttribute.JSONIgnore,
		IsIndexed:     rawAttribute.IsIndexed,
		IsMandatory:   rawAttribute.IsMandatory,
		IsReadOnly:    rawAttribute.ReadOnly,
		IsReference:   rawAttribute.IsReference,
		IsSensitive:   rawAttribute.SensibleData,
		LastModified:  time.Unix(0, rawAttribute.LastModifiedAt*int64(time.Millisecond)),
		Name:          rawAttribute.Name,
		RoleForUpdate: rawAttribute.RoleForUpdate,
		Type:          rawAttribute.Type,
		URL:           *aoms.MustUrl(url.Parse(rawAttribute.Href)),
		// Unexported
		isFromOtherModule: rawAttribute.AddedFromOtherModule,
		metaModelURL:      *aoms.MustUrl(url.Parse(rawAttribute.MetaModelHref)),
		refID:             rawAttribute.RefID,
	}
}
