//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/Jehanv60/app"
	"github.com/Jehanv60/controller"
	"github.com/Jehanv60/middleware"
	"github.com/Jehanv60/repository"
	"github.com/Jehanv60/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var barangSet = wire.NewSet(
	repository.NewRepositoryBarang,
	service.NewBarangService,
	controller.NewBarangController,
	wire.Bind(new(repository.BarangRepository), new(*repository.BarangRepoImpl)),
	wire.Bind(new(service.BarangService), new(*service.BarangServiceImpl)),
	wire.Bind(new(controller.BarangController), new(*controller.BarangControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDb,
		validator.New,
		barangSet,
		app.NewRouter,
		wire.Bind(
			new(http.Handler), new(*httprouter.Router),
		),
		middleware.NewAuthMiddleware,
		//NewServer,
	)
	return nil
}
