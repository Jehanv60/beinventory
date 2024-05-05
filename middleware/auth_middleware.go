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
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Mohon Login Dulu",
			}
			helper.WriteToResponse(w, webResponse)
		} else {
			switch r.URL.Path {
			case "/api/login":
				middleware.Handler.ServeHTTP(w, r)
			case "/api/pengguna":
				middleware.Handler.ServeHTTP(w, r)
			case "/api/logout":
				middleware.Handler.ServeHTTP(w, r)
			default:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				webResponse := web.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "Mohon Login Dulu",
				}
				helper.WriteToResponse(w, webResponse)
			}
		}
	} else if ApiKey == tokenn.Value {
		if r.Method == "POST" {
			switch r.URL.Path {
			case "/api/barang":
				middleware.Handler.ServeHTTP(w, r)
			case "/api/transaksi":
				middleware.Handler.ServeHTTP(w, r)
			default:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				webResponse := web.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "End Point Unauthorized",
				}
				helper.WriteToResponse(w, webResponse)
			}
		} else {
			middleware.Handler.ServeHTTP(w, r)
			tokenstring, err := util.Decodetoken(tokenn.Value)
			helper.PanicError(err)
			helper.WriteToResponse(w, map[string]interface{}{
				"User": tokenstring["pengguna"],
			})
		}
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
