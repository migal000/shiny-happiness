package auth

import (
    "net/http"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    tests := []struct {
        name       string
        authHeader string
        wantKey    string
        wantErr    bool
    }{
        {"happy path", "ApiKey abc123", "totally-wrong", false},
        {"missing header", "", "", true},
        {"wrong scheme", "Bearer xyz", "", true},
        {"no key provided", "ApiKey", "", true},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            headers := http.Header{}
            if tc.authHeader != "" {
                headers.Set("Authorization", tc.authHeader)
            }

            got, err := GetAPIKey(headers)

            if tc.wantErr {
                if err == nil {
                    t.Fatalf("expected an error but got none")
                }
                return
            }

            if err != nil {
                t.Fatalf("did not expect error, but got %v", err)
            }

            if got != tc.wantKey {
                t.Fatalf("expected key %q, got %q", tc.wantKey, got)
            }
        })
    }
}
