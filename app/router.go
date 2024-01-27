package app

import (
	"github.com/Jehanv60/controller"
	"github.com/Jehanv60/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(barangController controller.BarangController, penggunaController controller.PenggunaController) *httprouter.Router {
	//barang
	router := httprouter.New()
	router.GET("/api/barang", barangController.FindAll)
	router.GET("/api/barang/:barangId", barangController.FindById)
	router.PUT("/api/barang/:barangId", barangController.Update)
	router.DELETE("/api/barang/:barangId", barangController.Delete)
	router.POST("/api/barang", barangController.Create)
	//pengguna
	router.GET("/api/namapengguna/:NamaPengguna", penggunaController.FindByPenggunaRegister)
	router.GET("/api/pengguna", penggunaController.FindAll)
	router.GET("/api/pengguna/:penggunaId", penggunaController.FindById)
	router.PUT("/api/pengguna/:penggunaId", penggunaController.Update)
	router.POST("/api/pengguna", penggunaController.Create)
	router.POST("/api/login", penggunaController.LoginAuth)
	router.PanicHandler = exception.ErrorHandler
	return router
}
