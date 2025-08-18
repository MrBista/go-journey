package beforeaftertest

import (
	"fmt"
	"testing"

	"github.com/MrBista/go-journey/intermediate/15-testing/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	// hanya akan berjalan di package yang sama
	fmt.Println("Jalan sebelum semua test berjalan atau before all")

	m.Run()

	fmt.Println("Jalan setelah semua test berjalan atau after all")
}

func TestHelloBeforeAfter(t *testing.T) {
	res := utils.SayHello("Bisboy")

	assert.Equal(t, "Hello Bisboy", res)
}
