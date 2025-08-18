package goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func RunHellowWorld() {
	fmt.Println("Hello world")
}

func TestLearnBasicGoRutin(t *testing.T) {
	fmt.Println("Sebelum go rutine jalan")
	go RunHellowWorld() // ini akan berjalan async dan ga nunggu selesai function selanjutnya yang akan dipanggil
	fmt.Println("Setelah go rutine jalan")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Number ke " + strconv.Itoa(number))
}

func TestManyGorutine(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		go DisplayNumber(i)
	}
}
