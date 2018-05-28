package aomc_test

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/go-test/deep"

	"github.com/philippgille/apiomat-go/aomc"
	"github.com/philippgille/apiomat-go/aomc/dto"
)

// TestGetRawClasses tests if aomc.GetRawClasses leads to the correct aoms.Client call
// and the correct returned slice of structs
func TestGetRawClasses(t *testing.T) {
	// Prepare fake data and Get() implementation
	expected := []dto.Class{
		dto.Class{
			ID: "123",
		},
		dto.Class{
			ID: "456",
		},
	}
	expectedJsonBytes, err := json.Marshal(expected)
	stopOnError(err, t)
	expectedJson := string(expectedJsonBytes)
	moduleName := "fakeModule"
	expectedPath := "modules/" + moduleName + "/metamodels"
	onGet = func(path string, params url.Values) (string, error) {
		// Assertions
		if path != expectedPath {
			t.Errorf("path was %q, but should be %q", path, expectedPath)
		}
		if params != nil {
			t.Errorf("params was %v, but should be %v", params, nil)
		}
		// Assertions were okay, return fake data
		return expectedJson, nil
	}
	// Create fake client
	fakeClient := FakeClient{}
	client := aomc.NewClient(fakeClient)
	// Call method to test
	actual, err := client.GetRawClasses(moduleName)
	// Assertions
	stopOnError(err, t)
	if reflect.DeepEqual(actual, expected) == false {
		t.Errorf("actual %+v, expected %+v", actual, expected)
	}
}

// TestGetRawAttributes tests if aomc.GetRawAttributes leads to the correct aoms.Client call
// and the correct returned slice of structs
func TestGetRawAttributes(t *testing.T) {
	// Prepare fake data and Get() implementation
	expected := []dto.Attribute{
		dto.Attribute{
			ID: "123",
		},
		dto.Attribute{
			ID: "456",
		},
	}
	expectedJsonBytes, err := json.Marshal(expected)
	stopOnError(err, t)
	expectedJson := string(expectedJsonBytes)
	moduleName := "fakeModule"
	classId := "789"
	expectedPath := "modules/" + moduleName + "/metamodels/" + classId + "/attributes"
	onGet = func(path string, params url.Values) (string, error) {
		// Assertions
		if path != expectedPath {
			t.Errorf("path was %q, but should be %q", path, expectedPath)
		}
		if params != nil {
			t.Errorf("params was %v, but should be %v", params, nil)
		}
		// Assertions were okay, return fake data
		return expectedJson, nil
	}
	// Create fake client
	fakeClient := FakeClient{}
	client := aomc.NewClient(fakeClient)
	// Call method to test
	actual, err := client.GetRawAttributes(moduleName, classId)
	// Assertions
	stopOnError(err, t)
	if reflect.DeepEqual(actual, expected) == false {
		t.Errorf("actual %+v, expected %+v", actual, expected)
	}
}

// TestGetClasses tests if aomc.GetClasses leads to the correct aoms.Client call
// and the correct returned slice of structs
func TestGetClasses(t *testing.T) {
	// Prepare fake data and Get() implementation

	msec := time.Now().UnixNano() >> 6
	hrefString := "https://fake.url"
	expectedUrl, _ := url.Parse(hrefString)
	expected := []aomc.Class{
		// Don't fill all fields. Just all different types:
		// Slice of strings, aomc.Attribute, time, URL, string, bool
		aomc.Class{
			AllowedRolesCreate: []string{
				"someRole",
			},
			Attributes: []aomc.Attribute{
				// Don't fill the fields with much data, just check if the embedding works.
				// Attributes are tested more thoroughly in another test.
				aomc.Attribute{
					Id: "123",
					// Only needed because the empty JSON value is unmarshaled into the nil value for int, 0,
					// leading to a date from year 1970, while this struct leads to the nil value for time, which is year 0001.
					Created:      time.Unix(0, 0),
					LastModified: time.Unix(0, 0),
				},
				aomc.Attribute{
					Id: "456",
					// See above
					Created:      time.Unix(0, 0),
					LastModified: time.Unix(0, 0),
				},
			},
			// Only actual time
			Created:      time.Unix(0, msec*int64(time.Millisecond)),
			Href:         *expectedUrl,
			Id:           "789",
			IsDeprecated: true,
			// See above
			LastModified: time.Unix(0, 0),
		},
		aomc.Class{
			Attributes: []aomc.Attribute{},
			Id:         "101",
			// See above
			Created:      time.Unix(0, 0),
			LastModified: time.Unix(0, 0),
		},
	}
	expectedRawClasses := []dto.Class{
		dto.Class{
			AllowedRolesCreate: []string{
				"someRole",
			},
			CreatedAt:  msec,
			Href:       hrefString,
			ID:         "789",
			Deprecated: true,
		},
		dto.Class{
			ID: "101",
		},
	}
	expectedRawAttributes1 := []dto.Attribute{
		dto.Attribute{
			ID: "123",
		},
		dto.Attribute{
			ID: "456",
		},
	}
	expectedRawAttributes2 := []dto.Attribute{}
	expectedClassesJsonBytes, err := json.Marshal(expectedRawClasses)
	stopOnError(err, t)
	expectedAttributesJsonBytes1, err := json.Marshal(expectedRawAttributes1)
	stopOnError(err, t)
	expectedAttributesJsonBytes2, err := json.Marshal(expectedRawAttributes2)
	stopOnError(err, t)
	expectedClassesJson := string(expectedClassesJsonBytes)
	expectedAttributesJson1 := string(expectedAttributesJsonBytes1)
	expectedAttributesJson2 := string(expectedAttributesJsonBytes2)

	moduleName := "fakeModule"
	expectedPathClasses := "modules/" + moduleName + "/metamodels"
	expectedPathAttributes1 := "modules/" + moduleName + "/metamodels/789/attributes"
	expectedPathAttributes2 := "modules/" + moduleName + "/metamodels/101/attributes"
	validPaths := map[string]bool{
		expectedPathClasses:     true,
		expectedPathAttributes1: true,
		expectedPathAttributes2: true,
	}
	onGet = func(path string, params url.Values) (string, error) {
		var result string
		// Assertions
		if !validPaths[path] {
			t.Errorf("path was %v, but should be one of %v", path, validPaths)
		}
		if params != nil {
			t.Errorf("params was %v, but should be %v", params, nil)
		}
		// Assertions were okay, return fake data depending on the path
		switch path {
		case expectedPathClasses:
			result = expectedClassesJson
		case expectedPathAttributes1:
			result = expectedAttributesJson1
		case expectedPathAttributes2:
			result = expectedAttributesJson2
		}
		return result, nil
	}

	// Create fake client
	fakeClient := FakeClient{}
	client := aomc.NewClient(fakeClient)
	// Call method to test
	actual, err := client.GetClasses(moduleName)
	// Assertions
	stopOnError(err, t)
	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}

// TestGetAttributes tests if aomc.GetAttributes leads to the correct aoms.Client call
// and the correct returned slice of structs
func TestGetAttributes(t *testing.T) {
	// Prepare fake data and Get() implementation

	msec := time.Now().UnixNano() >> 6
	hrefString := "https://fake.url"
	expectedUrl, _ := url.Parse(hrefString)
	expected := []aomc.Attribute{
		// Don't fill all fields. Just all different types:
		// time, URL, string, bool
		aomc.Attribute{
			// Only actual time
			Created:  time.Unix(0, msec*int64(time.Millisecond)),
			Href:     *expectedUrl,
			Id:       "456",
			IsBinary: true,
			// Only needed because the empty JSON value is unmarshaled into the nil value for int, 0,
			// leading to a date from year 1970, while this struct leads to the nil value for time, which is year 0001.
			LastModified: time.Unix(0, 0),
		},
		aomc.Attribute{
			Id: "789",
			// See above
			Created:      time.Unix(0, 0),
			LastModified: time.Unix(0, 0),
		},
	}
	expectedRawAttributes := []dto.Attribute{
		dto.Attribute{
			CreatedAt: msec,
			Href:      hrefString,
			ID:        "456",
			Image:     true,
		},
		dto.Attribute{
			ID: "789",
		},
	}
	expectedJsonBytes, err := json.Marshal(expectedRawAttributes)
	stopOnError(err, t)
	expectedJson := string(expectedJsonBytes)

	moduleName := "fakeModule"
	expectedPath := "modules/" + moduleName + "/metamodels/123/attributes"
	onGet = func(path string, params url.Values) (string, error) {
		// Assertions
		if path != expectedPath {
			t.Errorf("path was %v, but should be %v", path, expectedPath)
		}
		if params != nil {
			t.Errorf("params was %v, but should be %v", params, nil)
		}
		// Assertions were okay, return fake data
		return expectedJson, nil
	}

	// Create fake client
	fakeClient := FakeClient{}
	client := aomc.NewClient(fakeClient)
	// Call method to test
	actual, err := client.GetAttributes(moduleName, "123")
	// Assertions
	stopOnError(err, t)
	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}
