package aomc_test

import (
	"net/url"
	"testing"
	"time"

	"github.com/go-test/deep"

	"github.com/philippgille/apiomat-go/aomc"
	"github.com/philippgille/apiomat-go/aomc/dto"
)

func TestConvertClassFromDto(t *testing.T) {
	// Prep
	// Not all fields must be tested, but one for each type.
	// Slice of strings, time, URL, string, bool
	msec := time.Now().UnixNano() >> 6
	hrefString := "https://fake.url"
	expectedUrl, _ := url.Parse(hrefString)
	expected := aomc.Class{
		AllowedRolesCreate: []string{
			"someRole",
		},
		// Only actual time
		Created:      time.Unix(0, msec*int64(time.Millisecond)),
		URL:          *expectedUrl,
		ID:           "123",
		IsDeprecated: true,
		// See above
		LastModified: time.Unix(0, 0),
	}
	given := dto.Class{
		AllowedRolesCreate: []string{
			"someRole",
		},
		CreatedAt:  msec,
		Href:       hrefString,
		ID:         "123",
		Deprecated: true,
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
	// time, URL, string, bool
	msec := time.Now().UnixNano() >> 6
	hrefString := "https://fake.url"
	expectedUrl, _ := url.Parse(hrefString)
	expected := aomc.Attribute{
		// Only actual time
		Created:      time.Unix(0, msec*int64(time.Millisecond)),
		URL:          *expectedUrl,
		ID:           "123",
		IsDeprecated: true,
		// See above
		LastModified: time.Unix(0, 0),
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
