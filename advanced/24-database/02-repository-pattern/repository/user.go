package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MrBista/go-journey/advanced/24-database/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	// FindAll(ctx context.Context) ([]*entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repository *userRepository) Create(ctx context.Context, user *entity.User) error {

	tx, err := repository.db.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {

		return fmt.Errorf("terjadi kesalahan saat membuka transaksi %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	script := "INSERT INTO users(username, password) VALUES(?, ?)"

	_, err = tx.ExecContext(ctx, script, user.Username, user.Password)

	if err != nil {
		return fmt.Errorf("terjadi kesalahan saat insert %w", err)
	}

	tx.Commit()

	return nil
}

func (repository *userRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {

	return nil, nil
}
