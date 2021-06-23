package config

import (
	admin "blog/admin/controllers"
	"blog/site/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Routes() *httprouter.Router  {
	r:= httprouter.New()

	// Admin index
	r.GET("/admin",admin.Dashboard{}.Index)

	// Article Operations
	r.GET("/admin/add-new-article",admin.Dashboard{}.AddNewArticle)
	r.POST("/admin/add-article",admin.Dashboard{}.AddArticle )
	r.GET("/admin/edit-article/:id",admin.Dashboard{}.UpdateArticleIndex)
	r.POST("/admin/update-article/:id",admin.Dashboard{}.UpdateArticle )
	r.GET("/admin/delete-article/:id",admin.Dashboard{}.DeleteArticle )

	// Category Operations
	r.GET("/admin/categories",admin.Category{}.Index)
	r.GET("/admin/add-new-category",admin.Category{}.AddNewCategory )
	r.POST("/admin/add-category",admin.Category{}.AddCategory )
	r.GET("/admin/edit-category/:id",admin.Category{}.UpdateCategoryIndex )
	r.POST("/admin/update-category/:id",admin.Category{}.UpdateCategory )
	r.GET("/admin/delete-category/:id",admin.Category{}.DeleteCategory )

	// Login Operations
	r.GET("/admin/login",admin.Login{}.Index)
	r.POST("/admin/do_login",admin.Login{}.Login)
	r.GET("/admin/logout",admin.Login{}.Logout)

	//SITE
	//Homepage
	r.GET("/",controllers.Homepage{}.Index)
	r.GET("/yazilar/:slug",controllers.Homepage{}.Detail)

	// SERVE FILES
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	r.ServeFiles("/assets/*filepath",http.Dir("site/assets"))
	return r
}
