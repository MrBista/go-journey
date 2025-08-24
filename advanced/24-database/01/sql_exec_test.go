package database

import (
	"context"
	"fmt"
	"testing"

	database "github.com/MrBista/go-journey/advanced/24-database"
)

// ddl

func TestInsertCustomer(t *testing.T) {
	db := database.GetConnection()

	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name) VALUES('BISBOY', 'GUSTI BISMAN')")

	if err != nil {
		// panic(err)
		t.Errorf("Terjadi keslahan saat insert ke customer table: %v", err)
	}
	fmt.Println("Success Insert Data to Database")
}
