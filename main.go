package main

import (
	"blog/admin/models"
	"blog/config"
	"net/http"
)

func main()  {
	models.User{}.Migrate()
	models.Article{}.Migrate()
	models.Category{}.Migrate()

	http.ListenAndServe(":8080",config.Routes())
}
