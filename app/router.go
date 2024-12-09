package app

import (
	"net/http"

	"github.com/Jehanv60/controller"
	"github.com/Jehanv60/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(barangController controller.BarangController, penggunaController controller.PenggunaController, transaksiController controller.TransaksiController) *httprouter.Router {
	//barang
	router := httprouter.New()
	router.GET("/api/barang", barangController.FindAll)
	router.GET("/api/barang/:barangId", barangController.FindById)
	router.PUT("/api/barang/:barangId", barangController.Update)
	router.DELETE("/api/barang/:barangId", barangController.Delete)
	router.POST("/api/barang", barangController.Create)
	//pengguna
	router.GET("/api/pengguna", penggunaController.FindAll)
	router.GET("/api/pengguna/:penggunaId", penggunaController.FindById)
	router.PUT("/api/pengguna/:penggunaId", penggunaController.Update)
	router.POST("/api/pengguna", penggunaController.Create)
	router.POST("/api/login", penggunaController.LoginAuth)
	router.POST("/api/logout", controller.Logout)
	//transaksi
	router.POST("/api/transaksi", transaksiController.Create)
	router.GET("/api/transaksi", transaksiController.ReportAll)
	router.NotFound = http.HandlerFunc(exception.NotFoundRouter())
	router.MethodNotAllowed = http.HandlerFunc(exception.MethodNotAllowed())
	router.PanicHandler = exception.ErrorHandler
	return router
}
