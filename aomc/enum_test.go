package aomc_test

import (
	"testing"

	"github.com/philippgille/apiomat-sdk-go/aomc"
)

func TestAuthImplStatus(t *testing.T) {
	expected := "UNKNOWN"
	given := aomc.Unknown
	actual := given.String()
	if expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestUserRole(t *testing.T) {
	expected := "User"
	given := aomc.User
	actual := given.String()
	if expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
