package models

import (
	"blog/admin/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName, LastName, UserName, Password, BirthDate string
}

func (user User) Migrate()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/User", "Migrate", "Database migration error", err.Error())
	}

	db.AutoMigrate(&user)
}

func (user User) Add()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/User", "Add", "Add article error", err.Error())
	}

	db.Create(&user)
}

func (user User) Get(where ...interface{}) User {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/User", "Get", "Get article error", err.Error())
		return user
	}

	db.First(&user, where...)
	return user
}

func (user User) GetAll(where ...interface{}) []User {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/User", "GetAll", "Get all articles error", err.Error())
		return nil
	}

	var users []User
	db.Find(&users, where...)
	return users
}

func (user User) Update(data User) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/User", "Update", "Update article error", err.Error())
	}

	db.Model(&user).Updates(data)
}

func (user User) Delete(){
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/User", "Update", "Update article error", err.Error())
	}

	db.Delete(&user)
}