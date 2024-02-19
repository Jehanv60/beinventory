package service

import (
	"context"
	"database/sql"

	"github.com/Jehanv60/exception"
	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/domain"
	"github.com/Jehanv60/model/web"
	"github.com/Jehanv60/repository"
	"github.com/Jehanv60/util"
	"github.com/go-playground/validator/v10"
)

type BarangServiceImpl struct {
	BarangRepository repository.BarangRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewBarangService(barangRepository repository.BarangRepository, DB *sql.DB, validate *validator.Validate) BarangService {
	return &BarangServiceImpl{
		BarangRepository: barangRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *BarangServiceImpl) Create(ctx context.Context, request web.BarangCreateRequest) web.BarangResponse {
	service.Validate.RegisterValidation("alphanumdash", util.ValidateSelf)
	err := service.Validate.Struct(request)
	util.ErrValidateSelf(err)
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs := domain.Barang{
		NameProd:   request.NameProd,
		Hargaprod:  request.Hargaprod,
		Keterangan: request.Keterangan,
		Stok:       request.Stok,
	}
	barangs = service.BarangRepository.Save(ctx, tx, barangs)
	return helper.ToBarangResponse(barangs)
}

func (service *BarangServiceImpl) Update(ctx context.Context, update web.BarangUpdate) web.BarangResponse {
	service.Validate.RegisterValidation("alphanumdash", util.ValidateSelf)
	err := service.Validate.Struct(update)
	util.ErrValidateSelf(err)
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs, err := service.BarangRepository.FindById(ctx, tx, update.Id)
	barangs.NameProd = update.NameProd
	barangs.Hargaprod = update.Hargaprod
	barangs.Keterangan = update.Keterangan
	barangs.Stok = update.Stok
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}
	barangs = service.BarangRepository.Update(ctx, tx, barangs)
	return helper.ToBarangResponse(barangs)
}

func (service *BarangServiceImpl) Delete(ctx context.Context, barangId int) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs, err := service.BarangRepository.FindById(ctx, tx, barangId)
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}
	service.BarangRepository.Delete(ctx, tx, barangs)
}

func (service *BarangServiceImpl) FindById(ctx context.Context, barangId int) web.BarangResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs, err := service.BarangRepository.FindById(ctx, tx, barangId)
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}
	return helper.ToBarangResponse(barangs)
}

func (service *BarangServiceImpl) FindAll(ctx context.Context) []web.BarangResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs := service.BarangRepository.FindAll(ctx, tx)
	return helper.ToBarangResponses(barangs)
}
