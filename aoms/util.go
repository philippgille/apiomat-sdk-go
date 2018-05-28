package aoms

import (
	"net/url"
)

// MustUrl enables parsing and dereferencing a URL in one line.
// It returns a panic if err is not nil.
// Call it for example when you want to dereference the pointer to a parsed URL
// and assign the value to a variable in one line of code.
// This can be useful when initializing structs for example.
//    var urlPtr *url.Url
//    var urlVal url.Url
//    // Assigning to a pointer is possible in one line
//    urlPtr, _ = url.Parse(urlString)
//    // But two lines are needed for the value
//    tempPtr, _ := url.Parse(urlString)
//    urlVal = *tempPtr
//
//    // With this method though:
//    urlVal = *aoms.MustUrl(url.Parse(urlString))
func MustUrl(url *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	return url
}
