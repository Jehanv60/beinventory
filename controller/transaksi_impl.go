package controller

import (
	"net/http"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/web"
	"github.com/Jehanv60/service"
	"github.com/Jehanv60/util"
	"github.com/julienschmidt/httprouter"
)

type TransaksiControllerImpl struct {
	TransactionService service.TransaksiService
	PenggunaService    service.PenggunaService
	BarangService      service.BarangService
}

func NewTransaksiController(transaksiService service.TransaksiService, penggunaService service.PenggunaService, barangService service.BarangService) TransaksiController {
	return &TransaksiControllerImpl{
		TransactionService: transaksiService,
		PenggunaService:    penggunaService,
		BarangService:      barangService,
	}
}

func (controller *TransaksiControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idUser := controller.PenggunaService.FindByPenggunaLogin(r.Context(), util.TokenEnv(r))
	transaksiCreateRequest := web.TransactionCreateRequest{}
	helper.ReadFromBody(r, &transaksiCreateRequest)
	transaksiResponse := controller.TransactionService.Create(r.Context(), transaksiCreateRequest, idUser.Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   transaksiResponse,
	}
	helper.WriteToResponse(w, webResponse)
}

func (controller *TransaksiControllerImpl) ReportAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idUser := controller.PenggunaService.FindByPenggunaLogin(r.Context(), util.TokenEnv(r))
	transaksiResponse := controller.TransactionService.ReportAll(r.Context(), idUser.Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   transaksiResponse,
	}
	helper.WriteToResponse(w, webResponse)
}
