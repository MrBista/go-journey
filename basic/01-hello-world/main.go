package main

import (
	"fmt"
	"go-journey/basic/01-hello-world/utils"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hai ini bisma menyapa semua orang")

	// import function to another packge
	utils.SayHello()
}
