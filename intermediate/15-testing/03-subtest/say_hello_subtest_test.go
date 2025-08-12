package subtest

import (
	"testing"

	"github.com/MrBista/go-journey/intermediate/15-testing/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSayHelloSubtest(t *testing.T) {
	t.Run("Bismen", func(t *testing.T) {
		result := utils.SayHello("Bismen")
		require.Equal(t, "Hello Bismen", result)
	})
	t.Run("Taka", func(t *testing.T) {
		result := utils.SayHello("Taka")
		require.Equal(t, "Hello Taka", result)
	})
}

// kegunaaan subtest bisa untuk dinamis test
func TestHelloSubtestTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hello Bismen",
			request:  "Bismen",
			expected: "Hello Bismen",
		},
		{
			name:     "Hello Taka",
			request:  "Taka",
			expected: "Hello Taka",
		},
		{
			name:     "Hello Gusti",
			request:  "Gusti",
			expected: "Hello Gusti",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := utils.SayHello(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
