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

/*
contoh penggunan tanpa interface
var DB = NewDb1()

	func NewDb1() *sql.DB {
		var err error
		var db *sql.DB
		db, err = sql.Open("postgres", "host=localhost user=han port=5432 password=solo dbname=pos1 sslmode=disable")
		if err != nil {
			panic(err)
		}
		db.SetMaxIdleConns(5)
		db.SetMaxOpenConns(20)
		db.SetConnMaxIdleTime(10 * time.Minute)
		db.SetConnMaxLifetime(60 * time.Minute)
		return db
	}

penggunaId, err := PenggunaSelect(penggunaCreateRequest.Pengguna)
helper.PanicError(err)

	func PenggunaSelect(NamaPengguna string) (domain.Pengguna, error) {
		var err error
		tx, err := DB.Begin()
		helper.PanicError(err)
		SQL := "select id, pengguna, email, password from pengguna where pengguna = $1"
		rows, err := tx.Query(SQL, NamaPengguna)
		helper.PanicError(err)
		pengguna := domain.Pengguna{}
		fmt.Println(pengguna)
		defer rows.Close()
		if rows.Next() {
			rows.Scan(&pengguna.Id, &pengguna.Pengguna, &pengguna.Email, &pengguna.Sandi)
			return pengguna, nil
		} else {
			return pengguna, nil
		}

}
*/
func (controller *PenggunaControllerImpl) LoginAuth(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	penggunaCreateRequest := web.LoginRequest{}
	helper.ReadFromBody(r, &penggunaCreateRequest)
	penggunaId := controller.PenggunaService.FindByPenggunaLogin(r.Context(), penggunaCreateRequest.Pengguna)
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
		claims["exp"] = time.Now().Add(time.Hour * 5).Unix()
		Token, err := util.GenerateToken(&claims)
		helper.PanicError(err)
		hehe := &http.Cookie{
			Name:     "token",
			Value:    Token,
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 5),
			HttpOnly: true,
		}
		http.SetCookie(w, hehe)
		helper.WriteToResponse(w, map[string]interface{}{
			"Message":  "Token Berhasil Dibuat",
			"Token":    Token,
			"Validasi": "Username Atau Email Dan Password Sesuai",
		})
	}
}
