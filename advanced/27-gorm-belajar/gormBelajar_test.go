package gormbelajar

import (
	"belajar-golang-gorm/models"
	"fmt"
	"strconv"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local" // bisa ambil config dari viper
	dsn := "bisma:bisma@tcp(127.0.0.1:4000)/main_database?charset=utf8mb4&parseTime=True&loc=Local" // bisa ambil config dari viper
	dialect := mysql.Open(dsn)

	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

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

func TestTransaction(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		user := models.User{
			Password: "rahasia-negara",
			Name: models.Name{
				FirstName:  "Mas",
				MiddleName: "Bisman",
				LastName:   "Baru",
			},
		}

		err := tx.Create(&user).Error

		if err != nil {
			return err
		}

		return nil

	})

	assert.Nil(t, err)
}

func TestQuerySingleObject(t *testing.T) {
	user := models.User{}
	result := db.First(&user)

	assert.Nil(t, result.Error)
	assert.Equal(t, 1, user.ID)

	user = models.User{}
	result = db.Take(&user, "id = ?", 5)
	assert.Nil(t, result.Error)
	assert.Equal(t, 5, user.ID)

	user = models.User{}
	result = db.Last(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, 15, user.ID)
}

func TestQueryInlineCondition(t *testing.T) {
	user := models.User{}
	result := db.Take(&user, "id = ?", 5)
	assert.Nil(t, result.Error)
	assert.Equal(t, 5, user.ID)
}

func TestQueryAllObject(t *testing.T) {
	user := []models.User{}

	// result := db.Where("id in ?", []string{"1", "2", "3", "4"}).Find(&user)

	result := db.Find(&user, "id in ?", []string{"1", "2", "3", "4"})
	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(user))
	// fmt.Println("user: ", user)
}

func TestQuerWhere(t *testing.T) {
	var users []models.User

	// kalau where lalu ada where lagi maka itu akan And
	result := db.
		Where("first_name like ?", "%mas%").
		Where("password like ?", "%rahasia%").
		Find(&users)
	assert.Nil(t, result.Error)
	fmt.Println(users)

	// kalau where lalu ada or maka kondisi akan jadi or
	users = []models.User{}

	result = db.
		Where("first_name like ?", "%mas%").
		Or("password like ?", "%rahasia%").
		Find(&users)

	assert.Nil(t, result.Error)
	fmt.Println("new users: ", users)
}

func TestNotQuer(t *testing.T) {
	var users []models.User

	result := db.Not("first_name like ?", "%nama ke%").
		Where("first_name like ?", "%gusti%").
		Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(users))
	fmt.Println(users)
}

func TestSelectField(t *testing.T) {
	var users []models.User

	result := db.Select("id", "first_name", "password").Find(&users)

	assert.Nil(t, result.Error)

	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
	}

}

func TestStructCondition(t *testing.T) {
	// bisa digunakan untuk dinasmis where
	// minusnya ga bisa ditambahkan or manual
	userCondition := models.User{
		Name: models.Name{
			FirstName: "Mas",
		},
	}
	var users []models.User
	result := db.Where(userCondition).Find(&users)

	assert.Nil(t, result.Error)
	fmt.Println(userCondition)
}

func TestMapCondition(t *testing.T) {
	// beda nya dengan struct condition adalah struct condition ga bisa kondisi zero value atau kosong

	mapCondition := map[string]interface{}{
		"middle_name": "",
	}

	var users []models.User

	result := db.Where(mapCondition).Find(&users)

	assert.Nil(t, result.Error)
	fmt.Println(users)
}

func TestOrderLimitOffest(t *testing.T) {
	var users []models.User
	result := db.Order("id asc, first_name asc").Limit(5).Offset(5).Find(&users)

	assert.Nil(t, result.Error)
	fmt.Println(users)
}

type UserResponse struct {
	ID        int
	Firstname string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse

	result := db.Model(&models.User{}).Select("id", "first_name", "last_name").Find(&users)

	assert.Nil(t, result.Error)
	fmt.Println(users)
}

func TestUpdate(t *testing.T) {
	user := models.User{}

	result := db.Take(&user, "id = ?", 1)

	assert.Nil(t, result.Error)

	user.Name.FirstName = "Berau"
	user.Password = "terubah"

	result = db.Save(&user)
	assert.Nil(t, result.Error)
}

func TestUpdateSelectedColumns(t *testing.T) {
	var user models.User
	findUser := db.Take(&user, "id = ?", 2)
	assert.Nil(t, findUser.Error)

	assert.NotNil(t, user.ID)

	// ini tuk satu column
	findUser.Update("password", "diubah passwordnya")

	assert.Nil(t, findUser.Error)
	assert.NotEqual(t, 0, findUser.RowsAffected)

	findUser.Updates(map[string]interface{}{
		"middle_name": "",
		"last_name":   "Bratha",
	})

	assert.Nil(t, findUser.Error)
	assert.NotEqual(t, 0, findUser.RowsAffected)

}
