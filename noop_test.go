package noop_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	noop "github.com/traefik-contrib/noop"
)

func TestNoop(t *testing.T) {
	testCases := []struct {
		desc        string
		cfg         *noop.Config
		expected    int
		expectedErr bool
	}{
		{
			desc:     "default config",
			cfg:      noop.CreateConfig(),
			expected: http.StatusTeapot,
		},
		{
			desc:     "custom config 200",
			cfg:      &noop.Config{ResponseCode: 200},
			expected: http.StatusOK,
		},
		{
			desc:        "custom config 1000 error",
			cfg:         &noop.Config{ResponseCode: 1000},
			expected:    1000,
			expectedErr: true,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

			handler, err := noop.New(ctx, next, test.cfg, "noop")
			if test.expectedErr {
				if err == nil {
					t.Fatal("error expectd")
				}
			} else {
				recorder := httptest.NewRecorder()

				req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
				if err != nil {
					t.Fatal(err)
				}

				handler.ServeHTTP(recorder, req)

				if recorder.Result().StatusCode != test.expected {
					t.Errorf("invalid response code: %d", recorder.Result().StatusCode)
				}
			}
		})
	}
}
