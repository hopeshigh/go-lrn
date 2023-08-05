package main

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestHandler_Fetch(t *testing.T) {
	mockClient := &MockClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("Hello, World!")),
			}, nil
		},
	}

	handler := NewHandler(mockClient)

	body, err := handler.Fetch("http://dummy-url.com")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(body) != "Hello, World!" {
		t.Errorf("got %v, want %v", string(body), "Hello, World!")
	}
}

func TestHandler_Decode(t *testing.T) {
	handler := NewHandler(http.DefaultClient)
	data := []byte(`{"userId": 1, "id": 1, "title": "test title", "body": "test body"}`)
	expected := &Post{
		UserId: 1,
		Id:     1,
		Title:  "test title",
		Body:   "test body",
	}
	post := &Post{}

	err := handler.Decode(data, post)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(post, expected) {
		t.Errorf("got %v, want %v", post, expected)
	}
}

func TestHandler_Transform(t *testing.T) {
	handler := NewHandler(http.DefaultClient)

	tests := []struct {
		name     string
		input    interface{}
		expected string
		isError  bool
	}{
		{
			name:     "valid post input",
			input:    &Post{Title: "abc"},
			expected: "nop",
			isError:  false,
		},
		{
			name:     "non-post type",
			input:    "I am a string",
			expected: "err",
			isError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := handler.Transform(tt.input)
			if tt.isError == false {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				} else if post, ok := tt.input.(*Post); ok && post.Title != tt.expected {
					t.Errorf("got %v, want %v", post.Title, tt.expected)
				}
			} else {
				if err == nil {
					t.Errorf("expected error, got %v", err)
				}
			}
		})
	}
}
