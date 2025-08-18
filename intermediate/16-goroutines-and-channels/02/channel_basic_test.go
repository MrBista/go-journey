package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestBasicChanel(t *testing.T) {
	// make(chan tipeData)

	// secara default chanel itu adalah pass by reference
	ch := make(chan string)

	fmt.Println("Sebelum go rutine")
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello world"
	}()

	fmt.Println("Setelah go rutine")

	data := <-ch

	fmt.Println("finish: data chanel dibalikan", data)

	fmt.Println(data)
	close(ch)
}

type Person struct {
	Name string
	Age  int
}

func GiveMeResponse(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "dari data yang dikelola function ini akan dikirimke chanel ini"
}

func ModifiedPerson(ch chan Person) {
	makePerson := Person{
		Name: "Bismen",
		Age:  10,
	}

	time.Sleep(2 * time.Second)
	ch <- makePerson
}

func TestChanelAsParameter(t *testing.T) {
	ch := make(chan string)

	chPerson := make(chan Person)

	go GiveMeResponse(ch)

	go ModifiedPerson(chPerson)

	fmt.Println("val: " + <-ch)
	fmt.Println(<-chPerson)
	close(ch)
	close(chPerson)
}
