package aomc_test

import (
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/philippgille/apiomat-go/aomc"
)

// TestGetRawBackends tests if aomc.TestGetRawBackends leads to the correct aomx.Client call
// and the correct returned slice of structs
func TestGetRawBackends(t *testing.T) {
	// Prepare fake Get() implementation
	expectedJsonBytes, err := ioutil.ReadFile("dto/application.json")
	stopOnError(err, t)
	expectedJson := string(expectedJsonBytes)
	customerName := "johnDoe"
	expectedPath := "customers/" + customerName + "/apps"
	onGet = func(path string, params url.Values) (string, error) {
		// Assertions
		if path != expectedPath {
			t.Errorf("path was %q, but should be %q", path, expectedPath)
		}
		if params != nil {
			t.Errorf("params was %v, but should be %v", params, nil)
		}
		// Assertions were okay, return fake data
		return "[" + expectedJson + "]", nil
	}
	// Create fake client
	fakeClient := FakeClient{}
	client := aomc.NewClient(fakeClient)
	// Call method to test
	backends, err := client.GetRawBackends(customerName)
	stopOnError(err, t)
	// Assertions
	if len(backends) != 1 {
		t.Errorf("expected %+v, actual %+v", 1, len(backends))
	}
	actual := backends[0]
	if actual.ApplicationName != "TestApp" {
		t.Errorf("expected %+v, actual %+v", "TestApp", actual.ApplicationName)
	}
	if actual.Configuration.LiveConfig["Basics"]["basics_createRole"] != "Guest" {
		t.Errorf("expected %+v, actual %+v", "Guest", actual.Configuration.LiveConfig["Basics"]["basics_createRole"])
	}
}

// TestGetRawBackendByName tests if aomc.TestGetRawBackendByName leads to the correct aomx.Client call
// and the correct returned structs
func TestGetRawBackendByName(t *testing.T) {
	// Prepare fake Get() implementation
	expectedJsonBytes, err := ioutil.ReadFile("dto/application.json")
	stopOnError(err, t)
	expectedJson := string(expectedJsonBytes)
	customerName := "johnDoe"
	appName := "TestApp"
	expectedPath := "customers/" + customerName + "/apps/" + appName
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
	actual, err := client.GetRawBackendByName(appName, customerName)
	stopOnError(err, t)
	// Assertions
	if actual.ApplicationName != "TestApp" {
		t.Errorf("expected %+v, actual %+v", "TestApp", actual.ApplicationName)
	}
	if actual.Configuration.LiveConfig["Basics"]["basics_createRole"] != "Guest" {
		t.Errorf("expected %+v, actual %+v", "Guest", actual.Configuration.LiveConfig["Basics"]["basics_createRole"])
	}
}
