package repositorypattern

import (
	"context"
	"fmt"
	"testing"

	database "github.com/MrBista/go-journey/advanced/24-database"
	"github.com/MrBista/go-journey/advanced/24-database/02-repository-pattern/repository"
	"github.com/MrBista/go-journey/advanced/24-database/entity"
)

func TestUserRepository(t *testing.T) {

	db := database.GetConnection()

	ctx := context.Background()

	userRepository := repository.NewUserRepository(db)

	userInsert := &entity.User{
		Username: "Bismen Coba",
		Password: "Bismen Password",
	}

	err := userRepository.Create(ctx, userInsert)

	if err != nil {
		t.Errorf("Terjadi kesalahan saat insert user %v", err)
	}

	fmt.Println("Berhasil insert user")

}
