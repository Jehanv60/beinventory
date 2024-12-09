package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

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
	apiKey := r.Header.Get(Header)
	tokenn, err := r.Cookie(Token)
	url := strings.TrimPrefix(r.URL.Path, "/api/")
	if apiKey == "" || err == http.ErrNoCookie {
		if r.Method == "POST" {
			switch url {
			case "login":
				middleware.Handler.ServeHTTP(w, r)
			case "pengguna":
				middleware.Handler.ServeHTTP(w, r)
			case "logout":
				if err == http.ErrNoCookie {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusNotFound)
					webResponse := web.WebResponse{
						Code:   http.StatusNotFound,
						Status: "Cookie Tidak Ditemukan",
					}
					helper.WriteToResponse(w, webResponse)
				} else {
					middleware.Handler.ServeHTTP(w, r)
				}
			default:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				webResponse := web.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "Mohon Untuk Login",
				}
				helper.WriteToResponse(w, webResponse)
			}
		} else {
			kata2 := []string{}
			if apiKey == "" {
				kata2 = append(kata2, "Header Key Tidak Ada")
			}
			if err == http.ErrNoCookie {
				kata2 = append(kata2, "Cookie Tidak Ditemukan")
			}
			kata2 = append(kata2, "Login Dulu Untuk Method Selain POST")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			webResponse := web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "Forbidden",
				Data:   kata2,
			}
			log.Println(kata2)
			helper.WriteToResponse(w, webResponse)
		}
	} else if apiKey == tokenn.Value {
		if r.Method == "POST" {
			switch url {
			case "barang":
				middleware.Handler.ServeHTTP(w, r)
			case "transaksi":
				middleware.Handler.ServeHTTP(w, r)
			default:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				webResponse := web.WebResponse{
					Code:   http.StatusNotFound,
					Status: "Not Found",
				}
				helper.WriteToResponse(w, webResponse)
			}
		} else {
			tokenstring, err := util.Decodetoken(apiKey)
			helper.PanicError(err)
			middleware.Handler.ServeHTTP(w, r)
			helper.WriteToResponse(w, map[string]interface{}{
				"User": tokenstring["pengguna"],
			})
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		_, err := util.Decodetoken(apiKey)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			helper.WriteToResponse(w, map[string]interface{}{
				"Code":    http.StatusBadRequest,
				"Status":  "Bad Request",
				"Message": err.Error(),
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
}
