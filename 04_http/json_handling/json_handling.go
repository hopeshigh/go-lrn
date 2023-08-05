package jsonhandling

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// {
//   "userId": 1,
//   "id": 1,
//   "title": "delectus aut autem",
//   "completed": false
// }

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Fetcher struct {
	Client HttpClient
}

func (f *Fetcher) FetchAndDecode(url string, v interface{}) error {
	// Parse the url
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	// Make the request
	response, err := f.Client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Check status code
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned %v", response.StatusCode)
	}

	// Decode JSON response
	dec := json.NewDecoder(response.Body)
	if err := dec.Decode(v); err != nil {
		return err
	}

	return nil
}
