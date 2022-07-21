package models

import (
	"github.com/jinzhu/gorm"
	"github.com/leonardodfelix/mysql-crud/pkg/config"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string `gorm:""json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("id = ?", id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(id int64) User {
	var user User
	db.Where("id = ?", id).Delete(user)
	return user
}
