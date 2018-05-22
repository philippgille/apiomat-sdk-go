package aomc

// RawClass struct generated with gojson
type RawClass struct {
	AllowedRolesCreate     []string `json:"allowedRolesCreate"`
	AllowedRolesGrant      []string `json:"allowedRolesGrant"`
	AllowedRolesRead       []string `json:"allowedRolesRead"`
	AllowedRolesWrite      []string `json:"allowedRolesWrite"`
	AttributesHref         string   `json:"attributesHref"`
	CreatedAt              int64    `json:"createdAt"`
	Deprecated             bool     `json:"deprecated"`
	ExtendsGeoModel        bool     `json:"extendsGeoModel"`
	Href                   string   `json:"href"`
	ID                     string   `json:"id"`
	IsGlobal               bool     `json:"isGlobal"`
	IsInvisible            bool     `json:"isInvisible"`
	IsTransient            bool     `json:"isTransient"`
	LastModifiedAt         int64    `json:"lastModifiedAt"`
	MethodsHref            string   `json:"methodsHref"`
	ModuleHref             string   `json:"moduleHref"`
	Name                   string   `json:"name"`
	RequiredUserRoleCreate string   `json:"requiredUserRoleCreate"`
	RequiredUserRoleGrant  string   `json:"requiredUserRoleGrant"`
	RequiredUserRoleRead   string   `json:"requiredUserRoleRead"`
	RequiredUserRoleWrite  string   `json:"requiredUserRoleWrite"`
	RestrictResourceAccess bool     `json:"restrictResourceAccess"`
	UseOwnAuth             string   `json:"useOwnAuth"`
}
