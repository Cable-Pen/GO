package main

import (

	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Email string `gorm:"unique"`
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func createUser(db *gorm.DB, name string, email string) (*User , error) {
	user := User{Name: name , Email: email}
	result := db.Create(&user)
	return &user , result.Error
}

func getUserByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	result := db.First(&user , id)
	return &user, result.Error
}

func getAllUser(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}

func updateuser(db *gorm.DB, id uint , name string, email string) error{
	var user User
	result := db.First(&user, id)
	if result.Error !=nil{
		return result.Error
	}

	user.Name = name
	user.Email = email
	db.Save(&user)
	return nil
}
func deleteuser(db *gorm.DB, id uint) error{
	var user User
	result := db.Delete(&user, id)
	return result.Error
}
func main() {
	dsn := "root:Atmzsdnb252546852@tcp(127.0.0.1:3306)/jojo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
		return
	}
	AutoMigrate(db)

	newUser, err := createUser(db, "aa", "aa@gmail.com")
	if err != nil {
		fmt.Println("Error creating user:", err)
		}else {
		fmt.Println("new user created:" , newUser)
	}
}
