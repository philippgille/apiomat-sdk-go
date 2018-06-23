package aomc_test

import (
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/go-test/deep"

	"github.com/philippgille/apiomat-sdk-go/aomc"
	"github.com/philippgille/apiomat-sdk-go/aomc/dto"
)

func TestConvertRawBackendsFromJSON(t *testing.T) {
	// Prep
	// Not all fields must be tested, but one for each type.
	// string, AnalyticsIds (represernting all structs with the three stages as fields), bool, Configuration (struct with three maps of maps of strings), int64

	adminName1 := "johnDoe"
	adminName2 := "jeanDoe"
	liveConfig := make(map[string]map[string]string)
	stagingConfig := make(map[string]map[string]string)
	liveConfig["Basics"] = make(map[string]string)
	stagingConfig["Basics"] = make(map[string]string)
	liveConfig["Basics"]["basics_createRole"] = "Guest"
	stagingConfig["Basics"]["basics_createRole"] = "Guest"
	nowMillis := time.Now().UnixNano() / int64(time.Millisecond)

	expected := []dto.Backend{
		dto.Backend{
			AdminName:       adminName1,
			AutoUpgradePlan: true,
			CreatedAt:       nowMillis,
		},
		dto.Backend{
			AdminName: adminName2,
		},
	}
	// Must be separate instructions because the structs aren't types
	expected[0].AnalyticsIds.Live = "01"
	expected[0].AnalyticsIds.Staging = "02"
	expected[0].Configuration.LiveConfig = liveConfig
	expected[0].Configuration.StagingConfig = stagingConfig

	given := `
		[
			{
				"adminName": "` + adminName1 + `",
				"analyticsIds": {
					"LIVE": "01",
					"STAGING": "02"
				},
				"autoUpgradePlan": true,
				"configuration": {
					"liveConfig": {
						"Basics": {
							"basics_createRole": "Guest"
						  }
					},
					"stagingConfig": {
						"Basics": {
							"basics_createRole": "Guest"
						  }
					}
				},
				"createdAt": ` + strconv.FormatInt(nowMillis, 10) + `
			},
			{
				"adminName": "` + adminName2 + `"
			}
		]`

	actual, err := aomc.ConvertRawBackendsFromJSON(given)
	stopOnError(err, t)
	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}

func TestConvertRawBackendFromJSON(t *testing.T) {
	// Prep
	// Not all fields must be tested, but one for each type.
	// string, AnalyticsIds (represernting all structs with the three stages as fields), bool, Configuration (struct with three maps of maps of strings), int64

	adminName1 := "johnDoe"
	liveConfig := make(map[string]map[string]string)
	stagingConfig := make(map[string]map[string]string)
	liveConfig["Basics"] = make(map[string]string)
	stagingConfig["Basics"] = make(map[string]string)
	liveConfig["Basics"]["basics_createRole"] = "Guest"
	stagingConfig["Basics"]["basics_createRole"] = "Guest"
	nowMillis := time.Now().UnixNano() / int64(time.Millisecond)

	expected := dto.Backend{
		AdminName:       adminName1,
		AutoUpgradePlan: true,
		CreatedAt:       nowMillis,
	}
	// Must be separate instructions because the structs aren't types
	expected.AnalyticsIds.Live = "01"
	expected.AnalyticsIds.Staging = "02"
	expected.Configuration.LiveConfig = liveConfig
	expected.Configuration.StagingConfig = stagingConfig

	given := `
		{
			"adminName": "` + adminName1 + `",
			"analyticsIds": {
				"LIVE": "01",
				"STAGING": "02"
			},
			"autoUpgradePlan": true,
			"configuration": {
				"liveConfig": {
					"Basics": {
						"basics_createRole": "Guest"
						}
				},
				"stagingConfig": {
					"Basics": {
						"basics_createRole": "Guest"
						}
				}
			},
			"createdAt": ` + strconv.FormatInt(nowMillis, 10) + `
		}`

	actual, err := aomc.ConvertRawBackendFromJSON(given)
	stopOnError(err, t)
	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}

func TestConvertRawClassesFromJSON(t *testing.T) {
	// Prep
	// Not all fields must be tested, but one for each type.
	// []string, string, int64, bool
	role1 := "someRole1"
	role2 := "someRole2"
	hrefString1 := "https://fake1.url"
	hrefString2 := "https://fake2.url"
	msec1 := time.Now().UnixNano() / int64(time.Millisecond)
	msec2 := time.Now().UnixNano() / int64(time.Millisecond)
	deprecated1 := true
	deprecated2 := false
	expected := []dto.Class{
		dto.Class{
			AllowedRolesCreate: []string{
				role1,
			},
			AttributesHref: hrefString1,
			CreatedAt:      msec1,
			Deprecated:     deprecated1,
		},
		dto.Class{
			AllowedRolesCreate: []string{
				role2,
			},
			AttributesHref: hrefString2,
			CreatedAt:      msec2,
			Deprecated:     deprecated2,
		},
	}
	given := `
		[
			{
				"createdAt":` + strconv.FormatInt(msec1, 10) + `,
				"allowedRolesCreate": ["` + role1 + `"],
				"deprecated": ` + strconv.FormatBool(deprecated1) + `,
				"attributesHref": "` + hrefString1 + `"
			},
			{
				"createdAt":` + strconv.FormatInt(msec2, 10) + `,
				"allowedRolesCreate": ["` + role2 + `"],
				"deprecated": ` + strconv.FormatBool(deprecated2) + `,
				"attributesHref": "` + hrefString2 + `"
			}
		]`

	// Test
	actual, err := aomc.ConvertRawClassesFromJSON(given)
	stopOnError(err, t)
	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}

func TestConvertRawClassFromJSON(t *testing.T) {
	// Prep
	// Not all fields must be tested, but one for each type.
	// []string, string, int64, bool
	role := "someRole1"
	hrefString := "https://fake1.url"
	msec := time.Now().UnixNano() / int64(time.Millisecond)
	deprecated := true
	expected := dto.Class{
		AllowedRolesCreate: []string{
			role,
		},
		AttributesHref: hrefString,
		CreatedAt:      msec,
		Deprecated:     deprecated,
	}
	given := `
		{
			"createdAt":` + strconv.FormatInt(msec, 10) + `,
			"allowedRolesCreate": ["` + role + `"],
			"deprecated": ` + strconv.FormatBool(deprecated) + `,
			"attributesHref": "` + hrefString + `"
		}`

	// Test
	actual, err := aomc.ConvertRawClassFromJSON(given)
	stopOnError(err, t)
	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}

func TestConvertRawAttributesFromJSON(t *testing.T) {
	// Prep
	// Not all fields must be tested, but one for each type.
	// bool, string, int64
	added1 := true
	added2 := false
	capName1 := "Foo"
	capName2 := "Bar"
	msec1 := time.Now().UnixNano() / int64(time.Millisecond)
	msec2 := time.Now().UnixNano() / int64(time.Millisecond)
	expected := []dto.Attribute{
		dto.Attribute{
			AddedFromOtherModule: added1,
			CapitalizedName:      capName1,
			CreatedAt:            msec1,
		},
		dto.Attribute{
			AddedFromOtherModule: added2,
			CapitalizedName:      capName2,
			CreatedAt:            msec2,
		},
	}
	given := `
	[
		{
			"createdAt":` + strconv.FormatInt(msec1, 10) + `,
			"addedFromOtherModule": ` + strconv.FormatBool(added1) + `,
			"capitalizedName": "` + capName1 + `"
		},
		{
			"createdAt":` + strconv.FormatInt(msec2, 10) + `,
			"addedFromOtherModule": ` + strconv.FormatBool(added2) + `,
			"capitalizedName": "` + capName2 + `"
		}
	]`

	// Test
	actual, err := aomc.ConvertRawAttributesFromJSON(given)
	stopOnError(err, t)
	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}

func TestConvertRawAttributeFromJSON(t *testing.T) {
	// Prep
	// Not all fields must be tested, but one for each type.
	// bool, string, int64
	added := true
	capName := "Foo"
	msec := time.Now().UnixNano() / int64(time.Millisecond)
	expected := dto.Attribute{
		AddedFromOtherModule: added,
		CapitalizedName:      capName,
		CreatedAt:            msec,
	}
	given := `
		{
			"createdAt":` + strconv.FormatInt(msec, 10) + `,
			"addedFromOtherModule": ` + strconv.FormatBool(added) + `,
			"capitalizedName": "` + capName + `"
		}`

	// Test
	actual, err := aomc.ConvertRawAttributeFromJSON(given)
	stopOnError(err, t)
	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}

func TestConvertClassFromDto(t *testing.T) {
	// Prep

	// Not all fields must be tested, but one for each type.
	// string, url.URL, time.Time, []string, AuthImplStatus, bool, UserRole
	// No Attributes field, because that's not part of the convert function.
	id := "123"
	hrefString := "https://fake.url"
	expectedUrl, _ := url.Parse(hrefString)
	msec := time.Now().UnixNano() / int64(time.Millisecond)
	allowedRoles := []string{
		"someRole",
	}
	authStatus := aomc.Unknown
	copyRoles := true
	userRole := aomc.User
	expected := aomc.Class{
		ID:  id,
		URL: *expectedUrl,
		// Only actual time
		Created:                    time.Unix(0, msec*int64(time.Millisecond)),
		AllowedRolesCreate:         allowedRoles,
		AuthImplStatus:             authStatus,
		IsCopyRolesToBinAttributes: copyRoles,
		RequiredRoleCreate:         userRole,
		// Only needed because the empty JSON value is unmarshaled into the nil value for int, 0,
		// leading to a date from year 1970, while this struct leads to the nil value for time, which is year 0001.
		LastModified: time.Unix(0, 0),
	}
	given := dto.Class{
		AllowedRolesCreate: allowedRoles,
		CreatedAt:          msec,
		Href:               hrefString,
		ID:                 id,
		RequiredUserRoleCreate: userRole.String(),
		RestrictResourceAccess: copyRoles,
		UseOwnAuth:             authStatus.String(), // if authStatus.String() is correct is tested in enum_test
	}

	// Test
	actual := aomc.ConvertClassFromDto(given)
	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}

func TestConvertAttributeFromDto(t *testing.T) {
	// Prep

	// Not all fields must be tested, but one for each type.
	// string, net.URL, time.Time, bool
	msec := time.Now().UnixNano() / int64(time.Millisecond)
	hrefString := "https://fake.url"
	expectedUrl, _ := url.Parse(hrefString)
	expected := aomc.Attribute{
		ID:  "123",
		URL: *expectedUrl,
		// Only actual time
		Created: time.Unix(0, msec*int64(time.Millisecond)),
		// Only needed because the empty JSON value is unmarshaled into the nil value for int, 0,
		// leading to a date from year 1970, while this struct leads to the nil value for time, which is year 0001.
		LastModified: time.Unix(0, 0),
		IsDeprecated: true,
	}
	given := dto.Attribute{
		CreatedAt:  msec,
		Href:       hrefString,
		ID:         "123",
		Deprecated: true,
	}

	// Test
	actual := aomc.ConvertAttributeFromDto(given)
	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}
