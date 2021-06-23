package models

import (
	"blog/admin/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title, Slug, Description, Content, PictureUrl string
	CategoryID int
}

func (article Article) Migrate()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Migrate", "Database migration error", err.Error())
	}

	db.AutoMigrate(&article)
}

func (article Article) Add()  {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Add", "Add article error", err.Error())
	}

	db.Create(&article)
}

func (article Article) Get(where ...interface{}) Article {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Get", "Get article error", err.Error())
		return article
	}

	db.First(&article, where...)
	return article
}

func (article Article) GetAll(where ...interface{}) []Article {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "GetAll", "Get all articles error", err.Error())
		return nil
	}

	var articles []Article
	db.Find(&articles, where...)
	return articles
}

func (article Article) Update(data Article) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Update", "Update article error", err.Error())
	}

	db.Model(&article).Updates(data)
}

func (article Article) Delete(){
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.LogJson("Error", "Models/Article", "Update", "Update article error", err.Error())
	}

	db.Delete(&article)
}