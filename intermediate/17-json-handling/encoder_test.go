package jsonhandling

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoderJson(t *testing.T) {
	t.Parallel()

	user := User{Name: "Bob", Age: 25}

	encoder := json.NewEncoder(os.Stdout).Encode(user)

	if encoder != nil {
		t.Fatalf("Failed to encode JSON: %v", encoder)
	}
	// Output: {"name":"Bob","age":25}

}
