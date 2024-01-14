package middleware

import (
	"net/http"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/web"
	"github.com/Jehanv60/util"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ApiKey := r.Header.Get("Authorization")
	tokenn, err := r.Cookie("token")
	if ApiKey == "" || err == http.ErrNoCookie {
		if r.Method != "POST" {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Mohon Login Dulu",
			}
			helper.WriteToResponse(w, webResponse)
		} else if r.Method == "POST" && r.URL.Path == "/api/login" {
			middleware.Handler.ServeHTTP(w, r)
		} else if r.Method == "POST" {
			middleware.Handler.ServeHTTP(w, r)
		}
	} else if ApiKey == tokenn.Value {
		tokenstring, err := util.Decodetoken(tokenn.Value)
		helper.PanicError(err)
		webResponse := web.WebResponse{
			Code:   200,
			Status: "Login Sukses dengan username",
			Data:   tokenstring["pengguna"],
		}
		helper.WriteToResponse(w, webResponse)
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helper.WriteToResponse(w, webResponse)
	}
}
