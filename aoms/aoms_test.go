package aoms_test

import (
	"testing"

	"github.com/philippgille/apiomat-go/aoms"
)

func TestSystem(t *testing.T) {
	cases := []struct {
		sys  aoms.System
		want string
	}{
		{aoms.Live, "LIVE"},
		{aoms.Staging, "STAGING"},
		{aoms.Test, "TEST"},
	}
	for _, c := range cases {
		got := c.sys.String()
		if got != c.want {
			t.Errorf("got %q, want %q", got, c.want)
		}
	}
}

func TestSystemCast(t *testing.T) {
	want := aoms.Live
	got := aoms.System("LIVE")
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
