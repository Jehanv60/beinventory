package exception

import (
	"net/http"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/web"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}
	if validationError(w, r, err) {
		return
	}
	internalServerError(w, r, err)
}
func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteToResponse(w, webResponse)
		return true
	} else {
		return false
	}

}
func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFound)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Data Tidak Ditemukan",
			Data:   exception.Error,
		}

		helper.WriteToResponse(w, webResponse)
		return true
	} else {
		return false
	}

}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server error",
		Data:   err,
	}

	helper.WriteToResponse(w, webResponse)
}
