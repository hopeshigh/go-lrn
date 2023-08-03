package filebuffer

import (
	"log"
	"os"
)

func WriteToFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	return err
}

func ReadFromFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}
	return data, err
}
