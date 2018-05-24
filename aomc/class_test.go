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

// TestGetRawAttributes tests if aomc.GetRawAttributes leads to the correct aoms.Client call and the correct returned JSON
func TestGetRawAttributes(t *testing.T) {
	// Prepare fake data and Get() implementation
	expected := []aomc.RawAttribute{
		aomc.RawAttribute{
			ID: "123",
		},
		aomc.RawAttribute{
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
