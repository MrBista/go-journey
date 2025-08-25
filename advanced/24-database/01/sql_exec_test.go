package database

import (
	"context"
	"database/sql"
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

func prepareInserUser(ctx context.Context, db *sql.DB) (*sql.Stmt, error) {
	script := "INSERT INTO users(username, password) VALUES(?, ?)"

	statment, err := db.PrepareContext(ctx, script)

	if err != nil {
		return nil, err
	}

	return statment, nil

}

func TestInsertSafety(t *testing.T) {
	db := database.GetConnection()

	// defer db.Close()

	defer func(db *sql.DB) {
		if db != nil {
			db.Close()
		}
	}(db)

	ctx := context.Background()

	insertUser, err := prepareInserUser(ctx, db)

	if err != nil {
		t.Errorf("Terjadi kesalahan saat insert user: %v", err)
	}

	if insertUser != nil {
		defer insertUser.Close()
	}

	username := "bismen_taka_2"
	password := "taka"

	result, err := insertUser.ExecContext(ctx, username, password)

	if err != nil {
		t.Errorf("Terjadi kesalahan saat insert data user %v", err)
	}

	val, err := result.RowsAffected()

	if err != nil {
		t.Errorf("No row effected %v", err)
	}

	fmt.Println("value inserted: ", val)

}
