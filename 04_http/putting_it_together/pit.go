package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const postsUrl = "https://jsonplaceholder.typicode.com/posts/1"

// {
// 	"userId": 1,
// 	"id": 1,
// 	"title": "words",
// 	"body:": "lots of words with line breaks"
// }

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func NewPost() *Post {
	return &Post{}
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Handler struct {
	Client HttpClient
}

func NewHandler(client HttpClient) *Handler {
	return &Handler{
		Client: client,
	}
}

func (h *Handler) Fetch(url string) ([]byte, error) {
	// Parse the url
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Make the request
	response, err := h.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Check status code
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned %v", response.StatusCode)
	}

	// Read the entire body
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
func (h *Handler) Decode(data []byte, v interface{}) error {
	// Decode json
	dec := json.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(v); err != nil {
		return err
	}

	return nil
}
func (h *Handler) Transform(data interface{}) error {
	switch v := data.(type) {
	case *Post:
		caesarCipher(&v.Title, 13)
	default:
		return fmt.Errorf("data is of unknown type: %v", v)
	}
	return nil
}
func (h *Handler) PrettyPrint(data interface{}) {
	switch v := data.(type) {
	case *Post:
		fmt.Printf("User Id: %v, Post Id: %v, \n Title: %v, \n Body: %v", v.UserId, v.Id, v.Title, v.Body)
	default:
		fmt.Printf("nothing to print for type %v", v)
	}
}

func caesarCipher(input *string, shift int) {
	var temp string
	for _, char := range *input {
		if 'a' <= char && char <= 'z' {
			temp += string('a' + (char-'a'+rune(shift))%26)
		} else if 'A' <= char && char <= 'Z' {
			temp += string('A' + (char-'a'+rune(shift))%26)
		} else {
			temp += string(char)
		}
	}

	*input = temp
}

func main() {
	np := NewPost()
	client := &http.Client{}
	handler := NewHandler(client)

	resp, err := handler.Fetch(postsUrl)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = handler.Decode(resp, np)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = handler.Transform(np)
	if err != nil {
		log.Fatalf("%v", err)
	}

	handler.PrettyPrint(np)
}
