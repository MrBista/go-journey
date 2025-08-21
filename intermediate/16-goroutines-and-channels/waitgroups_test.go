package goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func RunAscynchrounus(group *sync.WaitGroup) {
	defer group.Done()

	// code disini

	// group.Add(1) // kalau tambah disini dia ga akan jalan karena bisa aja dia concurent ke code selanjutnya saat .Wait() dipanggil dia dianggap akan kosong quew nya jadi akan lanjut code selanjutnya tanpa menunggu proses concurent dari waiting group

	fmt.Println("Hello")

	// time.Sleep(1 * time.Second)
}

func TestWaitingGroupBelajar(t *testing.T) {
	// waiting group itu kayak await kalau di js
	// digunakan untuk menunggu proses concurent lain selesai baru akan jalan

	group := &sync.WaitGroup{}

	group.Add(1)
	go RunAscynchrounus(group)

	// seperti async await
	group.Wait()
	// kalau ga pakai wait dia akan ga menunggu proses concurent di go rutine nya
	fmt.Println("Complete")
}
