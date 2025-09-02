package jsonhandling

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestDecoderJson(t *testing.T) {
	// --- IGNORE ---
	jsonStr := `{"name": "Alice", "age": 30}`
	reader := strings.NewReader(jsonStr)
	var user User
	err := json.NewDecoder(reader).Decode(&user)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}
	fmt.Println("user: ", user)
	if user.Name != "Alice" || user.Age != 30 {
		t.Errorf("Decoded user does not match expected values: got %+v", user)
	}
	// --- IGNORE ---
}
