package main

import (
	"net/http"
	"testing"

	"github.com/daniilgaltsev/learn-cicd-starter/internal/auth"
)
import _ "github.com/daniilgaltsev/learn-cicd-starter/internal/database"

func TestGetAPIKey(t *testing.T) {
	type testCase struct {
		name      string
		header    http.Header
		expected  string
		errNotNil bool
	}
	testCases := []testCase{
		{
			name:      "no auth header",
			header:    http.Header{},
			expected:  "",
			errNotNil: true,
		},
		{
			name: "malformed auth header",
			header: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expected:  "",
			errNotNil: true,
		},
		{
			name: "valid auth header",
			header: http.Header{
				"Authorization": []string{"ApiKey 123"},
			},
			expected:  "123",
			errNotNil: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := auth.GetAPIKey(tc.header)
			if got != tc.expected {
				t.Fatalf("expected `%s` got `%s`", tc.expected, got)
			}
			if tc.errNotNil && err == nil {
				t.Fatalf("expected an error, got `nil`")
			}
			if !tc.errNotNil && err != nil {
				t.Fatalf("expected no error, got `%s`", err)
			}
		})
	}
}
