package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func OnlyIn(ch chan<- string) {
	time.Sleep(2 * time.Second)

	// function ini memiliki parameter chanel yang hanya dapat mengirim chanel karena memiliki tipe only in yang ditandai dengan ch chan<- string
	ch <- "Data masuk ke chanel"
}

func OnlyOut(ch <-chan string) {
	time.Sleep(2 * time.Second)

	// function ini memiliki parameter chanel yang hanya dapat menerima chanel karena memiliki tipe only in yang ditandai dengan ch <-chan string
	fmt.Println(<-ch)
}

func TestInOutChanel(t *testing.T) {
	ch := make(chan string)

	go OnlyIn(ch)
	go OnlyOut(ch)

	time.Sleep(3 * time.Second)

	close(ch)
}
