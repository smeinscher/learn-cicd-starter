package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Errorf("no error presented when error should occur")
		return
	}
	headers.Add("Authorization", "ApiKey randomtoken")
	got, err := GetAPIKey(headers)
	want := "randomtoken"
	if err != nil {
		t.Errorf("error getting api key: %s", err)
		return
	}
	if got != want {
		t.Errorf("did not get what was expected")
		return
	}
}
