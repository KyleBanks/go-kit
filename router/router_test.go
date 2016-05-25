// Defines the Route interface, and registers routes to a server
package router

import (
	"net/http"
	"testing"
	"net/url"
)


func TestParam(t *testing.T) {
	// Ensure POST value is returned when both are set
	r := http.Request{
		PostForm: url.Values{
			"Test": []string{"POST"},
		},
		URL: &url.URL{
			RawQuery: "Test=GET",
		},
	}

	if "POST" != Param(&r, "Test") {
		t.Error("Expected Param() to return POST param when both POST and GET are set:", Param(&r, "Test"))
	}

	// Clear the POST, and ensure GET is returned
	r.PostForm = url.Values{}
	if "GET" != Param(&r, "Test") {
		t.Error("Expected Param() to return GET param when only GET is set:", Param(&r, "Test"))
	}

	// Clear both and ensure an empty string is returned
	r.URL = &url.URL{}
	if "" != Param(&r, "Test") {
		t.Error("Expected Param() to return empty string when neither are set:", Param(&r, "Test"))
	}
}