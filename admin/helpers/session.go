package helpers

import "net/http"

func CheckPageSession(w http.ResponseWriter, r *http.Request)  {
	if !CheckUserSession(w,r) {
		SetAlert(w,r,"Lütfen önce giriş yapınız!")
		http.Redirect(w,r,"/admin/login",http.StatusSeeOther)
		return
	}
}
