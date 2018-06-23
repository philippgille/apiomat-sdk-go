package aomx_test

import (
	"testing"

	"github.com/philippgille/apiomat-sdk-go/aomx"
)

func TestSystem(t *testing.T) {
	cases := []struct {
		sys  aomx.System
		want string
	}{
		{aomx.Live, "LIVE"},
		{aomx.Staging, "STAGING"},
		{aomx.Test, "TEST"},
	}
	for _, c := range cases {
		got := c.sys.String()
		if got != c.want {
			t.Errorf("got %q, want %q", got, c.want)
		}
	}
}

func TestSystemCast(t *testing.T) {
	want := aomx.Live
	got := aomx.System("LIVE")
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
