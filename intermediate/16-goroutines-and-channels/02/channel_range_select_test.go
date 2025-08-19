package goroutines

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRangeChannel(t *testing.T) {
	chanel := make(chan string)

	go func() {
		for i := 0; i < 100; i++ {
			chanel <- "Mengirim data ke " + strconv.Itoa(i)
			fmt.Println("Mengirim ", i)
		}
		close(chanel)
		fmt.Println("Chanel ditutup oleh produser")
	}()

	fmt.Println("Consumer mulai menerima")

	for v := range chanel {
		fmt.Println("Ditermia: ", v)
	}

	// for {
	// 	value, ok := <-chanel
	// 	if !ok {
	// 		fmt.Println("Channel sudah ditutup")
	// 		break
	// 	}
	// 	fmt.Printf("Ditermia: %v \n", value)
	// }
}

func TestSelectChanel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Mengirim data chanel pertama"
	}()

	go func() {
		ch2 <- "Mengirim data chanel kedua"
	}()

	counter := 0
	for {
		select {
		case data := <-ch1:
			fmt.Println("Data diterima dari Chanel 1", data)
			counter++
		case data := <-ch2:
			fmt.Println("Data ditermia dari Chanel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
