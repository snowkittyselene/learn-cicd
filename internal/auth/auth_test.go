package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input   http.Header
		wantStr string
		wantErr string
	}{
		"success":   {input: map[string][]string{"Authorization": {"ApiKey 12345"}}, wantStr: "12345", wantErr: ""},
		"no header": {input: map[string][]string{}, wantStr: "", wantErr: ErrNoAuthHeaderIncluded.Error()},
		"no ApiKey": {input: map[string][]string{"Authorization": {"12345 12345"}}, wantStr: "", wantErr: "malformed authorization header"},
		"no value":  {input: map[string][]string{"Authorization": {"ApiKey"}}, wantStr: "", wantErr: "malformed authorization header"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := GetAPIKey(tc.input)
			if err == nil {
				if res != tc.wantStr {
					t.Fatalf("expected: %s, got %s", tc.wantStr, res)
				}
			} else {
				if err.Error() != tc.wantErr {
					t.Fatalf("expected: %s, got %s", tc.wantErr, err.Error())
				}
			}
		})
	}
}
