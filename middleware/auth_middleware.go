package middleware

import (
	"net/http"
	"os"

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
	helper.GoDoEnv()
	var (
		Header = os.Getenv("Header")
		Token  = os.Getenv("Token")
	)
	ApiKey := r.Header.Get(Header)
	tokenn, err := r.Cookie(Token)
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
		middleware.Handler.ServeHTTP(w, r)
		tokenstring, err := util.Decodetoken(tokenn.Value)
		helper.PanicError(err)
		helper.WriteToResponse(w, map[string]interface{}{
			"User": tokenstring["pengguna"],
		})
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
