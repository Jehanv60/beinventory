package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/web"
	"github.com/Jehanv60/service"
	"github.com/Jehanv60/util"
	"github.com/julienschmidt/httprouter"
)

type BarangControllerImpl struct {
	BarangService   service.BarangService
	PenggunaService service.PenggunaService
}

func NewBarangController(barangService service.BarangService, penggunaService service.PenggunaService) BarangController {
	return &BarangControllerImpl{
		BarangService:   barangService,
		PenggunaService: penggunaService,
	}
}

func (controller *BarangControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idUser := controller.PenggunaService.FindByPenggunaLogin(r.Context(), util.TokenEnv(r))
	barangCreateRequest := web.BarangCreateRequest{}
	helper.ReadFromBody(r, &barangCreateRequest)
	barangResponse := controller.BarangService.Create(r.Context(), barangCreateRequest, idUser.Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   barangResponse,
	}
	helper.WriteToResponse(w, webResponse)
}

func (controller *BarangControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idUser := controller.PenggunaService.FindByPenggunaLogin(r.Context(), util.TokenEnv(r))
	barangUpdate := web.BarangUpdate{}
	helper.ReadFromBody(r, &barangUpdate)
	id, err := strconv.Atoi(params.ByName("barangId"))
	helper.PanicError(err)
	barangUpdate.Id = id
	barangUpdate.IdUser = idUser.Id
	barangResponse := controller.BarangService.Update(r.Context(), barangUpdate, idUser.Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   barangResponse,
	}
	helper.WriteToResponse(w, webResponse)
}

func (controller *BarangControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idUser := controller.PenggunaService.FindByPenggunaLogin(r.Context(), util.TokenEnv(r))
	id, err := strconv.Atoi(params.ByName("barangId"))
	helper.PanicError(err)
	controller.BarangService.Delete(r.Context(), id, idUser.Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicError(err)
}

func (controller *BarangControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idUser := controller.PenggunaService.FindByPenggunaLogin(r.Context(), util.TokenEnv(r))
	id, err := strconv.Atoi(params.ByName("barangId"))
	helper.PanicError(err)
	barangResponse := controller.BarangService.FindById(r.Context(), id, idUser.Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   barangResponse,
	}
	helper.WriteToResponse(w, webResponse)
}

func (controller *BarangControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idUser := controller.PenggunaService.FindByPenggunaLogin(r.Context(), util.TokenEnv(r))
	barangResponses := controller.BarangService.FindAll(r.Context(), idUser.Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   barangResponses,
	}
	helper.WriteToResponse(w, webResponse)
}
