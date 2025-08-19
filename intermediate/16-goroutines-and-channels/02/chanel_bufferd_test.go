package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestBufferedChannel(t *testing.T) {
	ch := make(chan string, 3)

	fmt.Println("Mengirim pesan 1")
	ch <- "Pesan 1"

	fmt.Println("Mengirim pesan 2")
	ch <- "Pesan 2"

	fmt.Println("Mengirim pesan 3")
	ch <- "Pesan 3"

	go func() {
		fmt.Println("Mengirim pesan 4")
		ch <- "Pesan 4"
		fmt.Println("Pesan 4 berakhir")
	}()

	time.Sleep(1 * time.Second)

	fmt.Printf("Chanel berisi %d pesan \n", len(ch))
	fmt.Printf("Kapasitas chanel %d \n", cap(ch))

	fmt.Println("Mulai menerima pesan")
	for i := 0; i < 4; i++ {
		msg := <-ch
		fmt.Println("Diterima", msg)

	}
	close(ch)
}
