package models

import "time"

// oleh gorm akan otomatis di binding kan secara otomatis, namun tidak disarankan untuk readability antar developer ygy
type User struct {
	ID        int       `gorm:"primary_key;column:id;autoIncrement"`
	Name      Name      `gorm:"embedded"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}

// secara default juga dia akan dibuat jamak dan snack case per kata, namun ada baiknya kita deklrasikan sendiri nama table dengan implement interface TableName
func (u *User) TableName() string {
	return "users"
}
