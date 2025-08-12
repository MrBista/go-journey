package assertion

import (
	"fmt"
	"testing"

	"github.com/MrBista/go-journey/intermediate/15-testing/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSayHelloAssertion(t *testing.T) {
	result := utils.SayHello("Bism")
	assert.Equal(t, "Hello Bismen", result) // this will call Fail
	fmt.Println("TestSayHelloAssertion finsih to call")
}

func TestSayHelloRequire(t *testing.T) {
	result := utils.SayHello("Kuja")
	require.Equal(t, "Hello Kujang", result) // this will call FailNow
	fmt.Println("TestSayHelloRequire finsih to call")

}
