package aomc

// RawAttribute struct generated with gojson
type RawAttribute struct {
	AddedFromOtherModule bool   `json:"addedFromOtherModule"`
	CapitalizedName      string `json:"capitalizedName"`
	CreatedAt            int64  `json:"createdAt"`
	Deprecated           bool   `json:"deprecated"`
	Href                 string `json:"href"`
	ID                   string `json:"id"`
	Image                bool   `json:"image"`
	IsCollection         bool   `json:"isCollection"`
	IsEmbeddedObject     bool   `json:"isEmbeddedObject"`
	IsIndexed            bool   `json:"isIndexed"`
	IsMandatory          bool   `json:"isMandatory"`
	IsReference          bool   `json:"isReference"`
	JSONIgnore           bool   `json:"jsonIgnore"`
	LastModifiedAt       int64  `json:"lastModifiedAt"`
	LowerCaseName        string `json:"lowerCaseName"`
	MetaModelHref        string `json:"metaModelHref"`
	Name                 string `json:"name"`
	ReadOnly             bool   `json:"readOnly"`
	RefID                string `json:"refId"`
	RoleForUpdate        string `json:"roleForUpdate"`
	SensibleData         bool   `json:"sensibleData"`
	Type                 string `json:"type"`
}
