package fileio

import (
	"fmt"
	"os"
	"testing"
)

func writeFileSimple() error {
	content := "Hello world \nAku menuliskan sebuah file untuk radio head"

	contentBytes := []byte(content)

	err := os.WriteFile("output_1.txt", contentBytes, 0644)

	if err != nil {
		return fmt.Errorf("Gagal menulis file: %v", err)
	}

	return nil
}

func TestWriteFileSimple(t *testing.T) {
	err := writeFileSimple()

	if err != nil {
		t.Fatalf("Terjadi kesalahan: %v", err)
	}
}
