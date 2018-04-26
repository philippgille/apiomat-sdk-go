package aoms

import (
	"testing"
)

func TestSystem(t *testing.T) {
	cases := []struct {
		sys  System
		want string
	}{
		{Live, "LIVE"},
		{Staging, "STAGING"},
		{Test, "TEST"},
	}
	for _, c := range cases {
		got := c.sys.String()
		if got != c.want {
			t.Errorf("got %q, want %q", got, c.want)
		}
	}
}
