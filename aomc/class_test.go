package aomc_test

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"

	"github.com/philippgille/apiomat-go/aomc"
)

// TestGetRawClasses tests if aomc.GetRawClasses leads to the correct aoms.Client call and the correct returned JSON
func TestGetRawClasses(t *testing.T) {
	// Prepare fake data and Get() implementation
	expected := []aomc.RawClass{
		aomc.RawClass{
			ID: "123",
		},
		aomc.RawClass{
			ID: "456",
		},
	}
	expectedJsonBytes, err := json.Marshal(expected)
	stopOnError(err, t)
	expectedJson := string(expectedJsonBytes)
	moduleName := "fakeModule"
	onGet = func(path string, params url.Values) (string, error) {
		// Assertions
		expectedPath := "modules/" + moduleName + "/metamodels"
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
