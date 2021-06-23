package controllers

import (
	"blog/admin/helpers"
	"blog/admin/log"
	"blog/admin/models"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type Category struct {}

func (category Category) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	helpers.CheckPageSession(w,r)

	view, err := template.ParseFiles(helpers.Include("category/list/")...)	// Buradaki `...` diziyi sanki string elemanları yan yana vermiş gibi gösterecek.
	if err != nil {
		log.LogJson("Error","Category","Index","Could not convert html files for go to read.",err.Error())
		return
	}

	data:= make(map[string]interface{})
	data["categories"] =models.Category{}.GetAll()
	data["alerts"] = helpers.GetAlert(w,r)
	view.ExecuteTemplate(w,"index",data)
}

func (category Category) AddNewCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	helpers.CheckPageSession(w,r)

	view, err := template.ParseFiles(helpers.Include("category/add/")...)	// Buradaki `...` diziyi sanki string elemanları yan yana vermiş gibi gösterecek.
	if err != nil {
		log.LogJson("Error","Category","Add","Could not convert html files for go to read.",err.Error())
		return
	}
	view.ExecuteTemplate(w,"index",nil)
}

func (category Category) AddCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	categoryName := r.FormValue("category-name")
	models.Category{
		Name: categoryName,
	}.Add()

	helpers.SetAlert(w,r,"Kayıt Başarılı!!")
	http.Redirect(w,r,"/admin/categories",http.StatusSeeOther)
}

func (category Category) DeleteCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	cat:= models.Category{}.Get(params.ByName("id"))
	cat.Delete()
	http.Redirect(w,r,"/admin",http.StatusSeeOther)
}

func (category Category) UpdateCategoryIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	helpers.CheckPageSession(w,r)

	view, err := template.ParseFiles(helpers.Include("category/edit/")...)	// Buradaki `...` diziyi sanki string elemanları yan yana vermiş gibi gösterecek.
	if err != nil {
		log.LogJson("Error","Category","Update","Could not convert html files for go to read.",err.Error())
		return
	}

	data:= make(map[string]interface{})
	data["category"] = models.Category{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w,"index",data)
}

func (category Category) UpdateCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	cat := models.Category{}.Get(params.ByName("id"))
	categoryName := r.FormValue("category-name")

	cat.Update(models.Category{
		Name: categoryName,
	})
	http.Redirect(w,r,"/admin/categories",http.StatusSeeOther)
}