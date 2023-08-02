package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		desc     string
		handler  http.HandlerFunc
		path     string
		expected string
	}{
		{
			desc:     "HelloHandler",
			handler:  helloHandler,
			path:     "/",
			expected: "Hello, World!",
		},
		{
			desc:     "GoodbyeHandler",
			handler:  goodbyeHandler,
			path:     "/goodbye",
			expected: "Goodbye, World!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodGet, tt.path, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			recorder := httptest.NewRecorder()
			tt.handler.ServeHTTP(recorder, request)

			if status := recorder.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			if recorder.Body.String() != tt.expected {
				t.Errorf("handler returned  unexpected body: got %v want %v", recorder.Body.String(), tt.expected)
			}
		})
	}
}
