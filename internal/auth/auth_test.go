package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)

	req.Header.Set("Authorization", "ApiKey 1234567890")

	got, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("Key test failed")
	}

	expected := "1234567890"

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected: %s, got: %s", expected, got)
	}
}
