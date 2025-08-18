package assertion

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/MrBista/go-journey/intermediate/15-testing/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSayHelloAssertion(t *testing.T) {
	result := utils.SayHello("Bismen")
	assert.Equal(t, "Hello Bismen", result) // this will call Fail
	fmt.Println("TestSayHelloAssertion finsih to call")
}

func TestSayHelloRequire(t *testing.T) {
	result := utils.SayHello("Kujang")
	require.Equal(t, "Hello Kujang", result) // this will call FailNow
	fmt.Println("TestSayHelloRequire finsih to call")

}

func TestDividedFunc(t *testing.T) {
	// result := utils.DividedFunc()
	test := []struct {
		val1     int
		val2     int
		name     string
		expected int
		wantErr  bool
	}{
		{
			val1:     3,
			val2:     0,
			name:     "test_zero",
			expected: 0,
			wantErr:  true,
		},
		{
			val1:     3,
			val2:     1,
			name:     "divided_normal",
			expected: 3 / 1,
			wantErr:  false,
		},
		{
			val1:     3,
			val2:     2,
			name:     "divided_normal",
			expected: 3 / 2,
			wantErr:  false,
		},
	}

	for i, v := range test {
		t.Run(v.name+strconv.Itoa(i), func(t *testing.T) {
			got, err := utils.DividedFunc(v.val1, v.val2)
			if v.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, v.expected, got)
		})
	}
}
