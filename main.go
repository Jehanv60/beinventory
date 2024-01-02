package main

import (
	_ "database/sql"
	"fmt"
	"net/http"

	"github.com/Jehanv60/app"
	"github.com/Jehanv60/controller"
	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/middleware"
	"github.com/Jehanv60/repository"
	"github.com/Jehanv60/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {
	DB := app.NewDb()
	validate := validator.New()
	barangRepository := repository.NewRepositoryBarang()
	barangService := service.NewBarangService(barangRepository, DB, validate)
	barangController := controller.NewBarangController(barangService)
	penggunaRepository := repository.NewRepositoryPengguna()
	penggunaService := service.NewPenggunaService(penggunaRepository, DB, validate)
	penggunaController := controller.NewPenggunaController(penggunaService)
	router := app.NewRouter(barangController, penggunaController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	fmt.Println("berhasil Konek")
	err := server.ListenAndServe()
	helper.PanicError(err)
}
