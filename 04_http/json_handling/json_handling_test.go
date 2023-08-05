package jsonhandling

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"
)

type MockClient struct {
	MockDo func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestFetchAndDecode(t *testing.T) {
	MockClient := &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte(`[{"userId": 1, "id": 1, "title": "Test Title", "completed": false}]`))),
			}, nil
		},
	}

	fetcher := &Fetcher{
		Client: MockClient,
	}

	var todos []Todo
	err := fetcher.FetchAndDecode("http://test.com", &todos)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	want := []Todo{
		{
			UserId:    1,
			Id:        1,
			Title:     "Test Title",
			Completed: false,
		},
	}
	if !reflect.DeepEqual(todos, want) {
		t.Errorf("expected %v, got %v", want, todos)
	}
}
