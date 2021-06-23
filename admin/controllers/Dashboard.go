package controllers

import (
	"blog/admin/helpers"
	"blog/admin/log"
	"blog/admin/models"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Dashboard struct {}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	helpers.CheckPageSession(w,r)

	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory":func(categoryId int) string {
			return models.Category{}.Get(categoryId).Name
		},
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
	}).ParseFiles(helpers.Include("dashboard/list/")...)
	if err != nil {
		log.LogJson("Error","Dashboard","Index","Could not convert html files for go to read.",err.Error())
		return
	}

	data:= make(map[string]interface{})
	data["articles"] =models.Article{}.GetAll()
	data["alerts"] = helpers.GetAlert(w,r)
	view.ExecuteTemplate(w,"index",data)
}

func (dashboard Dashboard) AddNewArticle(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	helpers.CheckPageSession(w,r)

	view, err := template.ParseFiles(helpers.Include("dashboard/add/")...)	// Buradaki `...` diziyi sanki string elemanları yan yana vermiş gibi gösterecek.
	if err != nil {
		log.LogJson("Error","Dashboard","Index","Could not convert html files for go to read.",err.Error())
		return
	}

	data:= make(map[string]interface{})
	data["categories"]=models.Category{}.GetAll()
	view.ExecuteTemplate(w,"index",data)
}

func (dashboard Dashboard) AddArticle(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	title := r.FormValue("article-title")
	description := r.FormValue("article-description")
	slug := slug.Make(title)
	content := r.FormValue("article-content")
	categoryId,_ := strconv.Atoi(r.FormValue("article-category"))

	// upload file
	r.ParseMultipartForm(10 << 20)
	file, header,_ := r.FormFile("article-picture")

	f,_ := os.OpenFile("uploads/"+header.Filename,os.O_WRONLY|os.O_CREATE, 0666)

	io.Copy(f,file)

	models.Article{
		Title: title,
		Slug: slug,
		Description: description,
		CategoryID: categoryId,
		Content: content,
		PictureUrl: "uploads/"+header.Filename,
	}.Add()

	helpers.SetAlert(w,r,"Kayıt Başarılı!!")
	http.Redirect(w,r,"/admin",http.StatusSeeOther)
}

func (dashboard Dashboard) DeleteArticle(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	article := models.Article{}.Get(params.ByName("id"))
	article.Delete()
	http.Redirect(w,r,"/admin",http.StatusSeeOther)
}

func (dashboard Dashboard) UpdateArticleIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	helpers.CheckPageSession(w,r)

	view, err := template.ParseFiles(helpers.Include("dashboard/edit/")...)	// Buradaki `...` diziyi sanki string elemanları yan yana vermiş gibi gösterecek.
	if err != nil {
		log.LogJson("Error","Dashboard","Index","Could not convert html files for go to read.",err.Error())
		return
	}

	data:= make(map[string]interface{})
	data["article"] = models.Article{}.Get(params.ByName("id"))
	data["categories"]=models.Category{}.GetAll()
	view.ExecuteTemplate(w,"index",data)
}

func (dashboard Dashboard) UpdateArticle(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	article := models.Article{}.Get(params.ByName("id"))
	title := r.FormValue("article-title")
	description := r.FormValue("article-description")
	slug := slug.Make(title)
	content := r.FormValue("article-content")
	categoryId,_ := strconv.Atoi(r.FormValue("article-category"))
	isSelected := r.FormValue("is_selected")

	var pictureUrl string

	if isSelected == "1" {
		// upload file
		r.ParseMultipartForm(10 << 20)
		file, header,_ := r.FormFile("article-picture")

		f,_ := os.OpenFile("uploads/"+header.Filename,os.O_WRONLY|os.O_CREATE, 0666)

		io.Copy(f,file)
		pictureUrl= "uploads/"+header.Filename
		os.Remove(article.PictureUrl)

	}else{
		pictureUrl = article.PictureUrl
	}

	article.Update(models.Article{
		Title: title,
		Slug: slug,
		Description: description,
		CategoryID: categoryId,
		Content: content,
		PictureUrl: pictureUrl,
	})

	http.Redirect(w,r,"/admin",http.StatusSeeOther)
}