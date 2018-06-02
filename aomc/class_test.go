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

// TestGetRawClasses tests if aomc.GetRawClasses leads to the correct aomx.Client call
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

// TestGetRawClassByName tests if aomc.GetRawClassByName leads to the correct aomx.Client call
// and the correct returned structs
func TestGetRawClassByName(t *testing.T) {
	// Prepare fake data and Get() implementation
	className := "Foo"
	expected := dto.Class{
		ID:   "123",
		Name: className,
	}
	dtosForExpectedJson := []dto.Class{
		expected,
		dto.Class{
			ID:   "456",
			Name: "Bar",
		},
	}
	expectedJsonBytes, err := json.Marshal(dtosForExpectedJson)
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
	actual, err := client.GetRawClassByName(moduleName, className)
	// Assertions
	stopOnError(err, t)
	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}

// TestGetClasses tests if aomc.GetClasses leads to the correct aomx.Client call
// and the correct returned slice of structs
func TestGetClasses(t *testing.T) {
	// Prepare fake data and Get() implementation

	msec := time.Now().UnixNano() / int64(time.Millisecond)
	hrefString := "https://fake.url"
	expectedUrl, _ := url.Parse(hrefString)
	expected := []aomc.Class{
		// Don't fill all fields. Just a few different types:
		// Slice of strings, aomc.Attribute, time, URL, string, bool
		// A complete set of types is tested in converter_test.
		aomc.Class{
			ID:      "789",
			URL:     *expectedUrl,
			Created: time.Unix(0, msec*int64(time.Millisecond)), // Only actual time
			// Only needed because the empty JSON value is unmarshaled into the nil value for int, 0,
			// leading to a date from year 1970, while this struct leads to the nil value for time, which is year 0001.
			LastModified: time.Unix(0, 0),
			AllowedRolesCreate: []string{
				"someRole",
			},
			Attributes: []aomc.Attribute{
				// Don't fill the fields with much data, just check if the embedding works.
				// Attributes are tested more thoroughly in another test.
				aomc.Attribute{
					ID:           "123",
					Created:      time.Unix(0, 0), // See above
					LastModified: time.Unix(0, 0), // See above
				},
				aomc.Attribute{
					ID:           "456",
					Created:      time.Unix(0, 0), // See above
					LastModified: time.Unix(0, 0), // See above
				},
			},
			IsDeprecated: true,
		},
		aomc.Class{
			ID:           "101",
			Created:      time.Unix(0, 0), // See above
			LastModified: time.Unix(0, 0), // See above
			Attributes:   []aomc.Attribute{},
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

// TestGetClassByName tests if aomc.GetClassByName leads to the correct aomx.Client call
// and the correct returned struct
func TestGetClassByName(t *testing.T) {
	// Prepare fake data and Get() implementation

	msec := time.Now().UnixNano() / int64(time.Millisecond)
	hrefString := "https://fake.url"
	expectedUrl, _ := url.Parse(hrefString)
	className := "Foo"
	expected := aomc.Class{
		ID:      "class1",
		URL:     *expectedUrl,
		Created: time.Unix(0, msec*int64(time.Millisecond)), // Only actual time
		// Only needed because the empty JSON value is unmarshaled into the nil value for int, 0,
		// leading to a date from year 1970, while this struct leads to the nil value for time, which is year 0001.
		LastModified: time.Unix(0, 0),
		AllowedRolesCreate: []string{
			"someRole",
		},
		Attributes: []aomc.Attribute{
			// Don't fill the fields with much data, just check if the embedding works.
			// Attributes are tested more thoroughly in another test.
			aomc.Attribute{
				ID:           "attribute1",
				Created:      time.Unix(0, 0), // See above
				LastModified: time.Unix(0, 0), // See above
			},
			aomc.Attribute{
				ID:           "attribute2",
				Created:      time.Unix(0, 0), // See above
				LastModified: time.Unix(0, 0), // See above
			},
		},
		IsDeprecated: true,
		Name:         className,
	}
	dtosForExpectedJson := []dto.Class{
		dto.Class{
			AllowedRolesCreate: []string{
				"someRole",
			},
			CreatedAt:  msec,
			Href:       hrefString,
			ID:         "class1",
			Deprecated: true,
			Name:       className,
		},
		dto.Class{
			ID:   "class2",
			Name: "Bar",
		},
	}
	expectedRawAttributes1 := []dto.Attribute{
		dto.Attribute{
			ID: "attribute1",
		},
		dto.Attribute{
			ID: "attribute2",
		},
	}
	expectedRawAttributes2 := []dto.Attribute{}
	expectedClassesJsonBytes, err := json.Marshal(dtosForExpectedJson)
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
	expectedPathAttributes1 := "modules/" + moduleName + "/metamodels/class1/attributes"
	expectedPathAttributes2 := "modules/" + moduleName + "/metamodels/class2/attributes"
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
	actual, err := client.GetClassByName(moduleName, className)
	// Assertions
	stopOnError(err, t)
	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}
