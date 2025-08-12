package main

import (
	"fmt"

	"github.com/MrBista/app-hello/helper"
)

func main() {
	sayHelloCall := helper.SayHello("Panggil dari package helper")
	fmt.Println("Hello " + sayHelloCall)
}
