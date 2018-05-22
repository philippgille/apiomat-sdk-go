package aomc_test

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"

	"github.com/philippgille/apiomat-go/aomc"
)

// TestGetClasses tests if aomc.GetRawClasses leads to the correct aoms.Client call and the correct returned JSON
func TestGetClasses(t *testing.T) {
	// Prepare fake data and Get() implementation
	want := []aomc.RawClass{
		aomc.RawClass{
			ID: "123",
		},
		aomc.RawClass{
			ID: "456",
		},
	}
	wantJsonBytes, err := json.Marshal(want)
	stopOnError(err, t)
	wantJson := string(wantJsonBytes)
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
		return wantJson, nil
	}
	// Create fake client
	fakeClient := FakeClient{}
	client := aomc.NewClient(fakeClient)
	// Call method to test
	got, err := client.GetRawClasses(moduleName)
	// Assertions
	stopOnError(err, t)
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
