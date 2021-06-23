package controllers

import (
	"blog/admin/helpers"
	"blog/admin/log"
	"blog/admin/models"
	"crypto/sha256"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type Login struct {}

func (login Login) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	view,err := template.ParseFiles(helpers.Include("login/")...)
	if err != nil {
		log.LogJson("Error","Login","Index","Could not convert html files for go to read.",err.Error())
		return
	}

	data:= make(map[string]interface{})
	data["alert"] = helpers.GetAlert(w,r)
	view.ExecuteTemplate(w,"index",data)
}

func (login Login) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	userName := r.FormValue("username")
	password := fmt.Sprintf("%x",sha256.Sum256([]byte(r.FormValue("password"))))

	user := models.User{}.Get("password = ? AND user_name = ?",password,userName)

	if user.UserName == userName && user.Password == password {
		helpers.SetUserSession(w,r,	userName,password)
		helpers.SetAlert(w,r,"Hoşgeldiniz Sn. "+user.FirstName+" "+user.LastName)
		http.Redirect(w,r,"/admin",http.StatusSeeOther)
	}else{
		helpers.SetAlert(w,r,"Kullanıcı adınız veya parolanız hatalı!")
		http.Redirect(w,r,"/admin/login",http.StatusSeeOther)
	}
}

func (login Login) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	helpers.RemoveUserSession(w,r)
	helpers.SetAlert(w,r,"Çıkış işlemi Başarılı!")
	http.Redirect(w,r,"/admin/login",http.StatusSeeOther)
}
