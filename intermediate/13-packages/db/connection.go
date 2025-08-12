package db

import "fmt"

var connection string

func init() {
	// init function akan langsung ke panggil jika package nya di import
	fmt.Println("Initializing dataabase  coneection")
	connection = "localhsot:5432"
}

func GetConnection() string {
	return connection
}
