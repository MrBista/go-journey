package main

import (
	"fmt"
	_ "go-journey/intermediate/13-packages/db"
	u "go-journey/intermediate/13-packages/utils"
)

func main() {
	fmt.Println("Hallo dunia package")

	u.CallPackageLearn()

	// db.GetConnection()

}
