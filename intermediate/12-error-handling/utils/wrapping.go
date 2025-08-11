package utils

import (
	"fmt"
	"os"
)

func readFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		// wrap error dengan %w dan akan tetap mengembalikan error aslinya
		return fmt.Errorf("Failed to read file %s: %w", filename, err)
	}

	return nil
}

func WrappingLearn() {
	err := readFile("bubu.txt")
	if err != nil {
		fmt.Println(err)
	}
}
