package controller

import (
	"net/http"
	"os"

	"github.com/Jehanv60/helper"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	helper.GoDoEnv()
	hehe := &http.Cookie{
		Name:     os.Getenv("Token"),
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, hehe)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	helper.WriteToResponse(w, map[string]interface{}{
		"Message": "Anda Sudah Logout",
	})
}
