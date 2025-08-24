package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	database "github.com/MrBista/go-journey/advanced/24-database"
	"github.com/MrBista/go-journey/advanced/24-database/entity"
)

func preaperStatmentSelectCustomer(ctx context.Context, db *sql.DB) *sql.Stmt {
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	// ! Di close di function yang mengunakan prepare statment ini
	// db := database.GetConnection()

	// defer db.Close()

	// ctx := context.Background()

	statement, err := db.PrepareContext(ctx, script)

	if err != nil {
		panic(err)
	}
	// defer statement.Close()

	return statement

}

func TestPrepareStatmentGet(t *testing.T) {
	db := database.GetConnection()

	defer db.Close()

	ctx := context.Background()

	getDataByStatment := preaperStatmentSelectCustomer(ctx, db)

	rows, err := getDataByStatment.QueryContext(ctx)

	if err != nil {
		t.Errorf("Terjadi kesalahan saat read context %v", err)
	}

	defer rows.Close()

	var listCustomer []entity.Customer
	for rows.Next() {
		customer := entity.Customer{}

		var email sql.NullString
		var birthDate sql.NullTime
		var createdAt sql.NullTime

		err := rows.Scan(&customer.Id, &customer.Name, &email, &customer.Balance, &customer.Rating, &birthDate, &customer.Married, &createdAt)

		if err != nil {
			t.Errorf("Terjadi kesalahan saat scan row %v", err)
		}

		if email.Valid {
			customer.Email = email.String
		}
		if birthDate.Valid {
			customer.BirthDate = birthDate.Time
		}
		if createdAt.Valid {
			customer.CreatedAt = createdAt.Time
		}

		listCustomer = append(listCustomer, customer)

	}

	result, err := json.Marshal(listCustomer)

	if err != nil {
		t.Errorf("Terjadi kesalahan : %v", err)

	}

	fmt.Println(string(result))
}
