package aomc_test

import (
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/go-test/deep"

	"github.com/philippgille/apiomat-go/aomc"
	"github.com/philippgille/apiomat-go/aomc/dto"
)

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
