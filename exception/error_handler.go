package exception

import (
	"log"
	"net/http"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/web"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}
	if sameFoundError(w, r, err) {
		return
	}
	if validationError(w, r, err) {
		return
	}
	internalServerError(w, r, err)
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
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

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
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

func sameFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(SameFound)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Status Unauthorized",
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
	log.Println(err)
	helper.WriteToResponse(w, webResponse)
}
