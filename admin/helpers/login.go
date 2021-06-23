package helpers

import (
	"blog/admin/models"
	"net/http"
)

func SetUserSession(w http.ResponseWriter, r *http.Request, username, password string) {
	session,_ := store.Get(r,"blog-user")
	session.Values["username"] = username
	session.Values["password"] = password

	session.Save(r,w)
}

func CheckUserSession(w http.ResponseWriter, r *http.Request) bool {
	session,_ := store.Get(r,"blog-user")

	username := session.Values["username"]
	password := session.Values["password"]

	user := models.User{}.Get("password = ? AND user_name = ?",password,username)
	if user.UserName == username && user.Password == password {
		return true
	}else{
		return false
	}
}

func RemoveUserSession(w http.ResponseWriter, r *http.Request) {
	session,_ := store.Get(r,"blog-user")
	session.Options.MaxAge = -1

	session.Save(r,w)
}