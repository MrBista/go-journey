package database

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	database "github.com/MrBista/go-journey/advanced/24-database"
	"github.com/MrBista/go-journey/advanced/24-database/entity"
)

// ddl

func TestQueryContext(t *testing.T) {
	db := database.GetConnection()

	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT id, name FROM customer")
	if err != nil {
		// panic(err)
		t.Errorf("Terjadi keslahan saat select ke customer table: %v", err)
	}
	defer rows.Close()

	var listCustomer []entity.Customer
	for rows.Next() {
		customer := entity.Customer{}

		err := rows.Scan(&customer.Id, &customer.Name)

		if err != nil {
			t.Errorf("Terjadi kesalahan saat scan row %v", err)
		}

		listCustomer = append(listCustomer, customer)

	}

	data, err := json.Marshal(listCustomer)

	if err != nil {
		t.Errorf("Terjadi kesalahan saat membuat format ke json %v", err)
	}

	fmt.Println(string(data))
}
