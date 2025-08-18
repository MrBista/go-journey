package benchmark

import (
	"testing"

	"github.com/MrBista/go-journey/intermediate/15-testing/utils"
)

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.SayHello("Bismen")
	}
}
