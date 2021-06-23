package controllers

import (
	"blog/admin/log"
	helpers2 "blog/site/helpers"
	"blog/site/models"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"time"
)

type Homepage struct {}

func (homepage Homepage) Index(w http.ResponseWriter,r *http.Request,params httprouter.Params)  {
	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory" : func(categoryId int) string {
			return models.Category{}.Get(categoryId).Name
		},
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
	}).ParseFiles(helpers2.Include("homepage/list/")...)

	if err != nil {
		log.LogJson("Error","Homepage","Index","Could not convert html files for go to read.",err.Error())
		return
	}

	data := make(map[string]interface{})
	data["articles"] = models.Article{}.GetAll()

	view.ExecuteTemplate(w,"index",data)
}

func (homepage Homepage) Detail(w http.ResponseWriter,r *http.Request,params httprouter.Params)  {
	view,err := template.ParseFiles(helpers2.Include("homepage/detail/")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["article"] = models.Article{}.Get("slug = ?",params.ByName("slug"))
	view.ExecuteTemplate(w,"index",data)
}