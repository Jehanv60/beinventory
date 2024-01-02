package middleware

import (
	"net/http"
	"strings"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/web"
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
	ApiKey := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(ApiKey) != 2 {
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

// func (middleware *AuthMiddleware) Login(w http.ResponseWriter, r *http.Request) {
// 	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
// 	if len(authHeader) != 2 {
// 		fmt.Println("Malformed token")
// 		w.WriteHeader(http.StatusUnauthorized)
// 		w.Write([]byte("Malformed Token"))
// 	} else {
// 		jwtToken := authHeader[1]
// 		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 			}
// 			return []byte(util.Secretkey), nil
// 		})

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			ctx := context.WithValue(r.Context(), "props", claims)
// 			// Access context values in handlers like this
// 			// props, _ := r.Context().Value("props").(jwt.MapClaims)
// 			middleware.Handler.ServeHTTP(w, r.WithContext(ctx))
// 		} else {
// 			fmt.Println(err)
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("Unauthorized"))
// 		}
// 	}

// }
