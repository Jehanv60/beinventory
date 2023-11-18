package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/web"
	"github.com/Jehanv60/service"
	"github.com/julienschmidt/httprouter"
)

type BarangControllerImpl struct {
	BarangService service.BarangService
}

func NewBarangController(barangService service.BarangService) BarangController {
	return &BarangControllerImpl{
		BarangService: barangService,
	}
}

func (controller *BarangControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	barangCreateRequest := web.BarangCreateRequest{}
	helper.ReadFromBody(r, &barangCreateRequest)
	barangResponse := controller.BarangService.Create(r.Context(), barangCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   barangResponse,
	}
	helper.WriteToResponse(w, webResponse)
}

func (controller *BarangControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	barangUpdate := web.BarangUpdate{}
	helper.ReadFromBody(r, &barangUpdate)
	id, err := strconv.Atoi(params.ByName("barangId"))
	helper.PanicError(err)
	barangUpdate.Id = id

	barangResponse := controller.BarangService.Update(r.Context(), barangUpdate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   barangResponse,
	}
	helper.WriteToResponse(w, webResponse)
}

func (controller *BarangControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("barangId"))
	helper.PanicError(err)
	controller.BarangService.Delete(r.Context(), id)
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
	id, err := strconv.Atoi(params.ByName("barangId"))
	helper.PanicError(err)
	barangResponse := controller.BarangService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   barangResponse,
	}
	helper.WriteToResponse(w, webResponse)
}

func (controller *BarangControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	barangResponses := controller.BarangService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   barangResponses,
	}
	helper.WriteToResponse(w, webResponse)
}
