package aomc_test

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/go-test/deep"

	"github.com/philippgille/apiomat-sdk-go/aomc"
	"github.com/philippgille/apiomat-sdk-go/aomc/dto"
)

// TestGetRawAttributes tests if aomc.GetRawAttributes leads to the correct aomx.Client call
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

// TestGetAttributes tests if aomc.GetAttributes leads to the correct aomx.Client call
// and the correct returned slice of structs
func TestGetAttributes(t *testing.T) {
	// Prepare fake data and Get() implementation

	msec := time.Now().UnixNano() / int64(time.Millisecond)
	hrefString := "https://fake.url"
	expectedUrl, _ := url.Parse(hrefString)
	expected := []aomc.Attribute{
		// Don't fill all fields. Just all different types:
		// string, net.URL, time.Time, bool
		aomc.Attribute{
			ID:  "456",
			URL: *expectedUrl,
			// Only actual time
			Created: time.Unix(0, msec*int64(time.Millisecond)),
			// Only needed because the empty JSON value is unmarshaled into the nil value for int, 0,
			// leading to a date from year 1970, while this struct leads to the nil value for time, which is year 0001.
			LastModified: time.Unix(0, 0),
			IsBinary:     true,
		},
		aomc.Attribute{
			ID:           "789",
			Created:      time.Unix(0, 0), // See above
			LastModified: time.Unix(0, 0), // See above
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
		return string(expectedJsonBytes), nil
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
