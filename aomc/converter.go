package aomc

import (
	"net/url"
	"time"

	"github.com/philippgille/apiomat-go/aomc/dto"
	"github.com/philippgille/apiomat-go/aoms"
)

// ConvertClassFromDto converts a raw class of type dto.Class into the convenience type aomc.Class.
// It does NOT fetch and assign the class attributes, it just converts existing fields.
func ConvertClassFromDto(rawClass dto.Class) Class {
	result := Class{
		// First go through all fields of the raw class that can be copied
		AllowedRolesCreate:     rawClass.AllowedRolesCreate,
		AllowedRolesGrant:      rawClass.AllowedRolesGrant,
		AllowedRolesRead:       rawClass.AllowedRolesRead,
		AllowedRolesWrite:      rawClass.AllowedRolesWrite,
		ID:                     rawClass.ID,
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
	result.attributesURL = *aoms.MustUrl(url.Parse(rawClass.AttributesHref))
	result.Created = time.Unix(0, rawClass.CreatedAt*int64(time.Millisecond))
	result.URL = *aoms.MustUrl(url.Parse(rawClass.Href))
	result.LastModified = time.Unix(0, rawClass.LastModifiedAt*int64(time.Millisecond))
	result.methodsURL = *aoms.MustUrl(url.Parse(rawClass.MethodsHref))
	result.moduleURL = *aoms.MustUrl(url.Parse(rawClass.ModuleHref))

	return result
}

// ConvertAttributeFromDto converts a raw attribute of type dto.Attribute into the convenience type aomc.Attribute.
func ConvertAttributeFromDto(rawAttribute dto.Attribute) Attribute {
	result := Attribute{
		// First go through all fields of the raw class that can be copied
		ID:                rawAttribute.ID,
		IsDeprecated:      rawAttribute.Deprecated,
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
	result.URL = *aoms.MustUrl(url.Parse(rawAttribute.Href))
	result.LastModified = time.Unix(0, rawAttribute.LastModifiedAt*int64(time.Millisecond))
	result.metaModelURL = *aoms.MustUrl(url.Parse(rawAttribute.MetaModelHref))

	return result
}
