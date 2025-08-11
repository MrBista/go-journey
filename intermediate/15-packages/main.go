package main

import (
	"fmt"
	_ "go-journey/intermediate/15-packages/db"
	u "go-journey/intermediate/15-packages/utils"
)

func main() {
	fmt.Println("Hallo dunia package")

	u.CallPackageLearn()

	// db.GetConnection()

}
