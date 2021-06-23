package helpers

import (
	"github.com/gorilla/sessions"
	"net/http"
)

// var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var store = sessions.NewCookieStore([]byte("user_name_key"))

func SetAlert(w http.ResponseWriter, r *http.Request, message string) {
	session,_ := store.Get(r,"go-alert")
	session.AddFlash(message)
	session.Save(r,w)
}

func GetAlert(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	session,_ := store.Get(r,"go-alert")

	data := make(map[string]interface{})
	flashes := session.Flashes()

	if len(flashes) > 0 {
		data["isAlert"] = true
		data["message"] = flashes[0]
	}else {
		data["isAlert"] = false
		data["message"] = nil
	}

	session.Save(r,w)
	return data
}