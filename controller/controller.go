package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BarangController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type PenggunaController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindByPenggunaRegister(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindByPenggunaLogin(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	LoginAuth(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
