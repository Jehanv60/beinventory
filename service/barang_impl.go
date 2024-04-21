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

func (service *BarangServiceImpl) Create(ctx context.Context, request web.BarangCreateRequest, idUser int) web.BarangResponse {
	service.Validate.RegisterValidation("alphanumdash", util.ValidateAlphanumdash)
	err := service.Validate.Struct(request)
	util.ErrValidateSelf(err)
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs := domain.Barang{
		Id:         idUser,
		IdUser:     idUser,
		KodeBarang: request.KodeBarang,
		NameProd:   request.NameProd,
		HargaProd:  request.HargaProd,
		JualProd:   request.JualProd,
		ProfitProd: request.JualProd - request.HargaProd,
		Keterangan: request.Keterangan,
		Stok:       request.Stok,
	}
	barangs, err = service.BarangRepository.Save(ctx, tx, barangs, idUser)
	if err != nil {
		panic(exception.NewNotEqual(err.Error()))
	}
	return helper.ToBarangResponse(barangs)
}

func (service *BarangServiceImpl) Update(ctx context.Context, update web.BarangUpdate, idUser int) web.BarangResponse {
	service.Validate.RegisterValidation("alphanumdash", util.ValidateAlphanumdash)
	err := service.Validate.Struct(update)
	util.ErrValidateSelf(err)
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs, err := service.BarangRepository.FindById(ctx, tx, update.Id, idUser)
	barangs.KodeBarang = update.KodeBarang
	barangs.NameProd = update.NameProd
	barangs.HargaProd = update.HargaProd
	barangs.JualProd = update.JualProd
	barangs.ProfitProd = barangs.JualProd - barangs.HargaProd
	barangs.Keterangan = update.Keterangan
	barangs.Stok = update.Stok
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}
	barangss, err := service.BarangRepository.FindByNameUpdate(ctx, tx, barangs.KodeBarang, barangs.NameProd, idUser)
	if barangs.Id == barangss.Id {
		barangss.HargaProd = update.HargaProd
		barangss.JualProd = update.JualProd
		barangss.ProfitProd = barangss.JualProd - barangss.HargaProd
		barangss.Keterangan = update.Keterangan
		barangss.Stok = update.Stok
		barangss, err = service.BarangRepository.Update(ctx, tx, barangs, idUser)
		if err != nil {
			panic(exception.NewNotEqual(err.Error()))
		}
		return helper.ToBarangResponse(barangss)
	}
	if err != nil {
		panic(exception.NewSameFound(err.Error()))
	}
	barangs, err = service.BarangRepository.Update(ctx, tx, barangs, idUser)
	if err != nil {
		panic(exception.NewNotEqual(err.Error()))
	}
	return helper.ToBarangResponse(barangs)
}

func (service *BarangServiceImpl) Delete(ctx context.Context, barangId int, idUser int) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs, err := service.BarangRepository.FindById(ctx, tx, barangId, idUser)
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}
	service.BarangRepository.Delete(ctx, tx, barangs, idUser)
}

func (service *BarangServiceImpl) FindById(ctx context.Context, barangId int, idUser int) web.BarangResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs, err := service.BarangRepository.FindById(ctx, tx, barangId, idUser)
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}
	return helper.ToBarangResponse(barangs)
}

func (service *BarangServiceImpl) FindByNameRegister(ctx context.Context, kodeBarang string, barangName string, idUser int) web.BarangResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs, err := service.BarangRepository.FindByNameRegister(ctx, tx, kodeBarang, barangName, idUser)
	if err != nil {
		panic(exception.NewSameFound(err.Error()))
	}
	return helper.ToBarangResponse(barangs)
}

func (service *BarangServiceImpl) FindByNameUpdate(ctx context.Context, kodeBarang string, barangName string, idUser int) web.BarangResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs, err := service.BarangRepository.FindByNameUpdate(ctx, tx, kodeBarang, barangName, idUser)
	if err != nil {
		panic(exception.NewSameFound(err.Error()))
	}
	return helper.ToBarangResponse(barangs)
}

func (service *BarangServiceImpl) FindAll(ctx context.Context, idUser int) []web.BarangResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs := service.BarangRepository.FindAll(ctx, tx, idUser)
	return helper.ToBarangResponses(barangs)
}
