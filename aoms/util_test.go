package aoms_test

import (
	"net/url"
	"testing"

	"github.com/philippgille/apiomat-go/aoms"
)

func TestMustUrl(t *testing.T) {
	// Test valid URL
	urlString := "https://fake.url"
	urlVal := aoms.MustUrl(url.Parse(urlString))
	if urlString != urlVal.String() {
		t.Errorf("expected %v, but was %v", urlString, *urlVal)
	}

	// Test invalid URL
	defer func() {
		if r := recover(); r == nil {
			t.Error("no panic occured, but was expected")
		}
	}()
	urlString = "://invalid.url" // While an empty string IS allowed (might be relative URL), an ":" at index 0 is NOT allowed
	urlVal = aoms.MustUrl(url.Parse(urlString))
}
