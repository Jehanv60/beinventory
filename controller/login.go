package controller

import (
	"net/http"
	"time"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/web"
	"github.com/Jehanv60/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func (controller *PenggunaControllerImpl) LoginAuth(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	penggunaCreateRequest := web.LoginRequest{}
	helper.ReadFromBody(r, &penggunaCreateRequest)
	penggunaId := controller.PenggunaService.FindByPengguna(r.Context(), penggunaCreateRequest.Pengguna)
	webResponse := web.LoginRequest{
		Pengguna: penggunaCreateRequest.Pengguna,
		Email:    penggunaCreateRequest.Email,
		Sandi:    penggunaCreateRequest.Sandi,
	}
	isvalid := util.Unhashpassword(webResponse.Sandi, penggunaId.Sandi)
	if webResponse.Pengguna == "" || webResponse.Email == "" || webResponse.Sandi == "" {
		helper.WriteToResponse(w, map[string]interface{}{
			"Code":    500,
			"Status":  "Bad Request",
			"Message": "Data Masih Kosong Mohon Dilengkapi",
		})
	} else if webResponse.Pengguna != penggunaId.Pengguna || webResponse.Email != penggunaId.Email || !isvalid {
		helper.WriteToResponse(w, map[string]interface{}{
			"Code":    500,
			"Status":  "Bad Request",
			"Message": "Username Atau Email Dan Password Tidak Sesuai",
		})
	} else {
		claims := jwt.MapClaims{}
		claims["pengguna"] = penggunaCreateRequest.Pengguna
		claims["email"] = penggunaCreateRequest.Email
		claims["sandi"] = penggunaCreateRequest.Sandi
		claims["exp"] = time.Now().Add(time.Second * 5).Unix()

		token, err := util.GenerateToken(&claims)
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			Expires:  time.Now().Add(time.Second * 5),
			HttpOnly: true,
		})
		// decodetoken, err := util.Decodetoken(tokern)
		// oketoken, _ := util.VerifyToken(tokern)
		helper.PanicError(err)
		helper.WriteToResponse(w, map[string]interface{}{
			"Message":  "Token Berhasil Dibuat",
			"Token":    token,
			"Validasi": "Username Atau Email Dan Password Sesuai",
		})
	}
}
