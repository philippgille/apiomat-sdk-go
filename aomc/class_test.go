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

	// CreateTestObjects doesn't create a class object with all fields, but with fields for all types.
	// A complete set of types is tested in converter_test.
	expectedClass1, expectedClassDto1, _, expectedAttributesJson1, err := CreateTestObjects("class1", "", "attribute1", "attribute2")
	stopOnError(err, t)
	expected := []aomc.Class{
		expectedClass1,
		aomc.Class{
			ID:           "class2",
			Created:      time.Unix(0, 0), // Required for reasons stated in comments of CreateTestObjects(...)
			LastModified: time.Unix(0, 0), // See above
			Attributes:   []aomc.Attribute{},
		},
	}
	expectedRawClasses := []dto.Class{
		expectedClassDto1,
		dto.Class{
			ID: "class2",
		},
	}
	expectedRawAttributes2 := []dto.Attribute{}
	expectedClassesJson, err := json.Marshal(expectedRawClasses)
	stopOnError(err, t)
	expectedAttributesJson2, err := json.Marshal(expectedRawAttributes2)
	stopOnError(err, t)

	module := "fakeModule"
	onGet = CreateTestOnGet(module, "class1", "class2", expectedClassesJson, expectedAttributesJson1, expectedAttributesJson2, t)

	// Create fake client
	fakeClient := FakeClient{}
	client := aomc.NewClient(fakeClient)
	// Call method to test
	actual, err := client.GetClasses(module)
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

	// CreateTestObjects doesn't create a class object with all fields, but with fields for all types.
	// A complete set of types is tested in converter_test.
	className := "Foo"
	expected, _, expectedClassJson, expectedAttributesJson, err := CreateTestObjects("class1", className, "attribute1", "attribute2")

	module := "fakeModule"
	expectedClassesJson := []byte("[" + string(expectedClassJson) + "]")
	onGet = CreateTestOnGet(module, "class1", "_", expectedClassesJson, expectedAttributesJson, nil, t)

	// Create fake client
	fakeClient := FakeClient{}
	client := aomc.NewClient(fakeClient)
	// Call method to test
	actual, err := client.GetClassByName(module, className)
	// Assertions
	stopOnError(err, t)
	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}

// CreateTestObjects creates an example class object with fields for all different types:
// Slice of strings, aomc.Attribute, time, URL, string, bool
// The return values are: The class object, the class DTO, the JSON of the class and the JSON of the attributes in the class.
func CreateTestObjects(classId string, className string, attrID1 string, attrID2 string) (aomc.Class, dto.Class, []byte, []byte, error) {
	msec := time.Now().UnixNano() / int64(time.Millisecond)
	hrefString := "https://fake.url"
	expectedUrl, _ := url.Parse(hrefString)

	// Class
	class := aomc.Class{
		ID:      classId,
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
				ID:           attrID1,
				Created:      time.Unix(0, 0), // See above
				LastModified: time.Unix(0, 0), // See above
			},
			aomc.Attribute{
				ID:           attrID2,
				Created:      time.Unix(0, 0), // See above
				LastModified: time.Unix(0, 0), // See above
			},
		},
		IsDeprecated: true,
		Name:         className,
	}

	// DTOs
	dtoClass := dto.Class{
		AllowedRolesCreate: []string{
			"someRole",
		},
		CreatedAt:  msec,
		Href:       hrefString,
		ID:         classId,
		Deprecated: true,
		Name:       className,
	}
	dtoAttributes := []dto.Attribute{
		dto.Attribute{
			ID: attrID1,
		},
		dto.Attribute{
			ID: attrID2,
		},
	}

	// JSON
	classJSON, err := json.Marshal(dtoClass)
	if err != nil {
		return aomc.Class{}, dto.Class{}, nil, nil, err
	}
	attributeJSON, err := json.Marshal(dtoAttributes)
	if err != nil {
		return aomc.Class{}, dto.Class{}, nil, nil, err
	}

	return class, dtoClass, classJSON, attributeJSON, nil
}

func CreateTestOnGet(module string, classID1 string, classID2 string, expectedClassesJson []byte, expectedAttributesJson1 []byte, expectedAttributesJson2 []byte, t *testing.T) func(path string, params url.Values) (string, error) {
	expectedPathClasses := "modules/" + module + "/metamodels"
	expectedPathAttributes1 := "modules/" + module + "/metamodels/" + classID1 + "/attributes"
	expectedPathAttributes2 := "modules/" + module + "/metamodels/" + classID2 + "/attributes"
	validPaths := map[string]bool{
		expectedPathClasses:     true,
		expectedPathAttributes1: true,
		expectedPathAttributes2: true,
	}
	return func(path string, params url.Values) (string, error) {
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
			result = string(expectedClassesJson)
		case expectedPathAttributes1:
			result = string(expectedAttributesJson1)
		case expectedPathAttributes2:
			result = string(expectedAttributesJson2)
		}
		return result, nil
	}
}
