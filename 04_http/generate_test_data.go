package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <url> <filename>")
		os.Exit(1)
	}

	url := os.Args[1]
	filename := os.Args[2]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error making request to %s: %v", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("error creating file %s: %v", filename, err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("error writing to file: %v", err)
	}

	fmt.Println("Successfully wrote response to", filename)
}
