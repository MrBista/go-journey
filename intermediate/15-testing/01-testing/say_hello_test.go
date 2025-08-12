package testing

import (
	"fmt"
	"testing"

	"github.com/MrBista/go-journey/intermediate/15-testing/utils"
)

func TestSayHello(t *testing.T) {
	result := utils.SayHello("Besme")

	if result != "Hello Bismen" {
		// panic("must Hello Bismen") // bukan bestpractie
		t.Fail() // code setelahnya akan tetap jalan
	}

	fmt.Println("TestSayHello kelar")
}

func TestSayHelloTaka(t *testing.T) {
	result := utils.SayHello("Tak")

	if result != "Hello Taka" {
		t.FailNow() // akan langsung berhenti
	}
	fmt.Println("TestSayHelloTaka kelar")
}

func TestSayHelloError(t *testing.T) {
	result := utils.SayHello("Bisme")

	if result != "Hello Bismen" {
		// panic("must Hello Bismen") // bukan bestpractie
		t.Error("Expected must be Hello Bismen") // code setelahnya akan tetap jalan -> mirip dengan Fail
	}

	fmt.Println("TestSayHello kelar")
}

func TestSayHelloTakaFatal(t *testing.T) {
	result := utils.SayHello("Tak")

	if result != "Hello Taka" {
		t.Fatal("Expected must be Hello Taka") // akan langsung berhenti mirip dengan FaileNow
	}
	fmt.Println("TestSayHelloTaka kelar")
}
