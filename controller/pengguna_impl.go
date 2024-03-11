package controller

import (
	"net/http"
	"strconv"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/web"
	"github.com/Jehanv60/service"
	"github.com/julienschmidt/httprouter"
)

type PenggunaControllerImpl struct {
	PenggunaService service.PenggunaService
}

func NewPenggunaController(penggunaService service.PenggunaService) PenggunaController {
	return &PenggunaControllerImpl{
		PenggunaService: penggunaService,
	}
}

// Create implements PenggunaController.
func (controller *PenggunaControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	penggunaCreateRequest := web.PenggunaCreateRequest{}
	helper.ReadFromBody(r, &penggunaCreateRequest)
	controller.PenggunaService.FindByPenggunaRegister(r.Context(), penggunaCreateRequest.Pengguna, penggunaCreateRequest.Email)
	penggunaResponse := controller.PenggunaService.Create(r.Context(), penggunaCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   penggunaResponse,
	}
	helper.WriteToResponse(w, map[string]interface{}{
		"Code":   webResponse.Code,
		"Status": webResponse.Status,
		"Mesage": "Data Berhasil Ditambahkan",
	})
}

// FindAll implements PenggunaController.
func (controller *PenggunaControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	penggunaResponse := controller.PenggunaService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   penggunaResponse,
	}
	helper.WriteToResponse(w, webResponse)
}

// FindById implements PenggunaController.
func (controller *PenggunaControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("penggunaId"))
	helper.PanicError(err)
	penggunaResponse := controller.PenggunaService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   penggunaResponse,
	}
	helper.WriteToResponse(w, webResponse)
}

// FindById implements PenggunaController.
func (controller *PenggunaControllerImpl) FindByPenggunaRegister(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("namaPengguna")
	email := params.ByName("email")
	penggunaResponse := controller.PenggunaService.FindByPenggunaRegister(r.Context(), id, email)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   penggunaResponse,
	}
	helper.WriteToResponse(w, webResponse)
}

// FindById implements PenggunaController.
func (controller *PenggunaControllerImpl) FindByPenggunaLogin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("namaPengguna")
	penggunaResponse := controller.PenggunaService.FindByPenggunaLogin(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   penggunaResponse,
	}
	helper.WriteToResponse(w, webResponse)
}

// Update implements PenggunaController.
func (controller *PenggunaControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	penggunaUpdate := web.PenggunaUpdate{}
	helper.ReadFromBody(r, &penggunaUpdate)
	id, err := strconv.Atoi(params.ByName("penggunaId"))
	helper.PanicError(err)
	penggunaUpdate.Id = id
	controller.PenggunaService.FindByPenggunaRegister(r.Context(), penggunaUpdate.Pengguna, penggunaUpdate.Email)
	penggunaResponse := controller.PenggunaService.Update(r.Context(), penggunaUpdate)
	helper.PanicError(err)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   penggunaResponse,
	}
	helper.WriteToResponse(w, map[string]interface{}{
		"Code":   webResponse.Code,
		"Status": webResponse.Status,
		"Mesage": "Data Berhasil Diupdate",
	})
}
