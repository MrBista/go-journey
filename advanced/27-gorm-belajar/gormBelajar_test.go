package gormbelajar

import (
	"belajar-golang-gorm/models"
	"strconv"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local" // bisa ambil config dari viper
	dsn := "bisma:bisma@tcp(127.0.0.1:4000)/main_database?charset=utf8mb4&parseTime=True&loc=Local" // bisa ambil config dari viper
	dialect := mysql.Open(dsn)

	db, err := gorm.Open(dialect, &gorm.Config{})

	if err != nil {

		panic(err)
	}

	return db
}

var db = OpenConnection()

func TestConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("insert into sample(id, name) values(?, ?)", "1", "Eko").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?, ?)", "2", "Budi").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?, ?)", "3", "Kurniawan").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?, ?)", "4", "Joko").Error
	assert.Nil(t, err)
}

type Sample struct {
	Id   string
	Name string
}

func TestRawSql(t *testing.T) {
	var sample Sample
	err := db.Raw("SELECT id, name from sample where id = ?", "1").Scan(&sample).Error

	assert.Nil(t, err)

	assert.Equal(t, "1", sample.Id)

	var samples []Sample
	err = db.Raw("SELECT id, name from sample").Scan(&samples).Error

	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))

	logrus.Info(sample)
	logrus.Info(samples)

}

func TestCreateUser(t *testing.T) {
	user := models.User{
		Password: "rahasia",
		Name: models.Name{
			FirstName:  "Gustii",
			MiddleName: "Bisman",
			LastName:   "Taka",
		},
	}

	response := db.Create(&user)

	err := response.Error

	assert.Nil(t, err)

	assert.Equal(t, 1, int(response.RowsAffected))

	user.ID = 0
}

func TestBatchInsert(t *testing.T) {
	var users []models.User

	for i := 0; i < 10; i++ {
		user := models.User{
			Name: models.Name{
				FirstName: "Nama ke " + strconv.Itoa(i+1),
			},
			Password: "rahasia",
		}

		users = append(users, user)
	}

	result := db.Create(users)
	err := result.Error
	assert.Nil(t, err)

	assert.Equal(t, 10, int(result.RowsAffected))
}
