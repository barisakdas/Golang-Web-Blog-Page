package models

import (
	"blog/admin/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string
}

func (category Category) Migrate()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Migrate", "Database migration error", err.Error())
	}

	db.AutoMigrate(&category)
}

func (category Category) Add()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Add", "Add article error", err.Error())
	}

	db.Create(&category)
}

func (category Category) Get(where ...interface{}) Category {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Get", "Get article error", err.Error())
		return category
	}

	db.First(&category, where...)
	return category
}

func (category Category) GetAll(where ...interface{}) []Category {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "GetAll", "Get all articles error", err.Error())
		return nil
	}

	var categories []Category
	db.Find(&categories, where...)
	return categories
}

func (category Category) Update(data Category) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Update", "Update article error", err.Error())
	}

	db.Model(&category).Updates(data)
}

func (category Category) Delete(){
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Update", "Update article error", err.Error())
	}

	db.Delete(&category)
}