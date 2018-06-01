package aomc_test

import (
	"net/url"
	"testing"
)

// FakeClient is a fake implementation of the aomx.Client interface.
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

func stopOnError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("got error %v", err)
	}
}
