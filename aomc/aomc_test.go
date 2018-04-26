package aomc_test

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"

	"github.com/philippgille/apiomat-go/aomc"

	"github.com/philippgille/apiomat-go/aoms"
)

// FakeClient is a fake implementation of the aoms.Client interface.
// Its methods call functions that can be defined by each test on its own.
type FakeClient struct{}

func (c FakeClient) GetVersion() (string, error) {
	return onGetVersion()
}

func (c FakeClient) Get(path string, params url.Values) (string, error) {
	return onGet(path, params)
}

// Functions called by the fake client's methods, can be defined by each test
var onGetVersion = func() (string, error) {
	return "1.2.3", nil
}
var onGet func(string, url.Values) (string, error)

// TestGetClasses tests if aomc.GetClasses leads to the correct aoms.Client call and the correct returned JSON
func TestGetClasses(t *testing.T) {
	// Prepare fake data and Get() implementation
	want := []aomc.Class{
		aomc.Class{
			ID: "123",
		},
		aomc.Class{
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
	got, err := client.GetClasses(moduleName, aoms.Live)
	// Assertions
	stopOnError(err, t)
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func stopOnError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("got error %v", err)
	}
}
