package exception

import (
	"log"
	"net/http"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/web"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, err) {
		return
	}
	if sameNotEqual(w, err) {
		return
	}
	if sameFoundError(w, err) {
		return
	}
	if validationError(w, err) {
		return
	}
	internalServerError(w, err)
}

func validationError(w http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(ValidateFound)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var result []string
		for _, errVal := range exception.Error {
			result = append(result, errVal.Error())
		}
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   result,
		}
		helper.WriteToResponse(w, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(NotFound)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}
		helper.WriteToResponse(w, webResponse)
		return true
	} else {
		return false
	}
}

func sameFoundError(w http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(SameFound)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		webResponse := web.WebResponse{
			Code:   http.StatusFound,
			Status: "Status Found",
			Data:   exception.Error,
		}
		helper.WriteToResponse(w, webResponse)
		return true
	} else {
		return false
	}
}

func sameNotEqual(w http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(NotEqual)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error,
		}
		helper.WriteToResponse(w, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server error",
		Data:   err,
	}
	log.Println(err)
	helper.WriteToResponse(w, webResponse)
}
