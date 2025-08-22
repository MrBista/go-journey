package fileio

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

// region read file by os
func readFileExampleTxt() ([]byte, error) {
	content, err := os.ReadFile("example.txt") // untuk read file

	if err != nil {
		return nil, fmt.Errorf("Gagal untuk mendapatkan file: %v", err)
	}

	return content, nil
}
func TestReadFileLearn(t *testing.T) {
	content, err := readFileExampleTxt()

	if err != nil {
		t.Fatalf("Gagal read file example.txt karena %v", err)
	}

	fmt.Println(string(content))
}

// endregion

// region read file open first
func readFileWitOpenFirst() ([]byte, error) {
	file, err := os.Open("example.txt")

	if err != nil {
		return nil, fmt.Errorf("Gagal membuka file : %v", err)
	}

	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		return nil, fmt.Errorf("Gagal membaca file : %v", err)
	}

	return content, nil
}

func TestReadFileOpenFirst(t *testing.T) {
	content, err := readFileWitOpenFirst()

	if err != nil {
		t.Fatalf("Gagal mendapatkan file karena: %v", err)
	}

	fmt.Println(string(content))
}

// endregion

// region read with bufio

func readFileWithBufio() ([]byte, error) {
	file, err := os.Open("example.txt")

	if err != nil {
		return nil, fmt.Errorf("Gagal membuka file : %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var buffer bytes.Buffer

	for scanner.Scan() {
		buffer.Write(scanner.Bytes())
		buffer.WriteByte('\n')
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error saat membaca file: %v", err)
	}

	content := buffer.Bytes()
	if len(content) > 0 && content[len(content)-1] == '\n' {
		content = content[:len(content)-1]
	}

	return content, nil

}

func TestReadFileBufio(t *testing.T) {
	content, err := readFileWithBufio()

	if err != nil {
		t.Fatalf("Gagal mendapatkan file karena: %v", err)
	}

	fmt.Println(string(content))
}
