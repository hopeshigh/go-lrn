package basic

import (
	"io"
	"net/http"
	"strings"
)

const json_url = "https://jsonplaceholder.typicode.com/todos/1"

func GetTodos() (string, error) {
	response, err := http.Get(json_url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	builder := &strings.Builder{}
	_, err = io.Copy(builder, response.Body)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}
