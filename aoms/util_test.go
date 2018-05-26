package aoms_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/philippgille/apiomat-go/aoms"
)

func TestConvertUnixMillisToTime(t *testing.T) {
	// Prep
	now := time.Now()
	// Add 123 milliseconds and set local time zone
	loc, _ := time.LoadLocation("Local")
	expected := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 123000, loc)

	// Test the objects themselves
	msec := expected.UnixNano() >> 3
	actual := aoms.ConvertUnixMillisToTime(msec)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %+v, but was %+v", expected, actual)
	}

	// Also test their string representations,
	// just to make sure the objects weren't constructed in a wrong way and two nonsense objects are equal.
	expectedString := fmt.Sprint(expected)
	actualString := fmt.Sprint(actual)
	if expectedString != actualString {
		t.Errorf("expected %v, but was %v", expectedString, actualString)
	}

	// Also test the year,
	// in case the first test only tested two nonsense objects for equality
	// and the second test two weird strings (like year "634987")
	if now.Year() != actual.Year() {
		t.Errorf("expected %v, but was %v", now.Year(), actual.Year())
	}
}
