package service

import (
	"context"
	"database/sql"
	"fmt"

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
	barang := service.BarangRepository.FindByNameRegister(ctx, tx, request.KodeBarang, request.NameProd, idUser)
	if barang.KodeBarang == request.KodeBarang {
		panic(exception.NewSameFound(fmt.Sprintf("%s Sudah Digunakan, Mohon Untuk Cek Di Inventory", request.KodeBarang)))
	}
	if barang.NameProd == request.NameProd {
		panic(exception.NewSameFound(fmt.Sprintf("%s Sudah Digunakan, Mohon Untuk Cek Di Inventory", request.NameProd)))
	}
	barangs := domain.Barang{
		IdUser:     idUser,
		KodeBarang: request.KodeBarang,
		NameProd:   request.NameProd,
		HargaProd:  request.HargaProd,
		JualProd:   request.JualProd,
		ProfitProd: request.JualProd - request.HargaProd,
		Keterangan: request.Keterangan,
		Stok:       request.Stok,
	}
	barangs = service.BarangRepository.Save(ctx, tx, barangs, idUser)
	//Validasi jika input dirubah dari inspect elemen
	if barangs.ProfitProd < 0 {
		panic(exception.NewNotEqual("Harga Profit Tidak Boleh Dibawah Harga Jual"))
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
	barangs := service.BarangRepository.FindById(ctx, tx, update.Id, idUser)
	if barangs.Id != update.Id {
		panic(exception.NewNotFound("Data Barang Tidak Ditemukan"))
	}
	barangs.KodeBarang = update.KodeBarang
	barangs.NameProd = update.NameProd
	barangs.HargaProd = update.HargaProd
	barangs.JualProd = update.JualProd
	barangs.ProfitProd = barangs.JualProd - barangs.HargaProd
	barangs.Keterangan = update.Keterangan
	barangs.Stok = update.Stok
	barangss := service.BarangRepository.FindByNameUpdate(ctx, tx, barangs.KodeBarang, barangs.NameProd, idUser)
	if barangs.Id == barangss.Id {
		barangs = service.BarangRepository.Update(ctx, tx, barangs, idUser)
		if barangs.ProfitProd < 0 {
			panic(exception.NewNotEqual("Harga Profit Tidak Boleh Dibawah Harga Jual"))
		}
		return helper.ToBarangResponse(barangs)
	}
	if barangs.Id != barangss.Id {
		if barangs.KodeBarang == barangss.KodeBarang {
			panic(exception.NewSameFound(fmt.Sprintf("%s Sudah Digunakan, Mohon Untuk Cek Di Inventory", update.KodeBarang)))
		}
		if barangs.NameProd == barangss.NameProd {
			panic(exception.NewSameFound(fmt.Sprintf("%s Sudah Digunakan, Mohon Untuk Cek Di Inventory", update.NameProd)))
		}
	}
	barangs = service.BarangRepository.Update(ctx, tx, barangs, idUser)
	if barangs.ProfitProd < 0 {
		panic(exception.NewNotEqual("Harga Profit Tidak Boleh Dibawah Harga Jual"))
	}
	return helper.ToBarangResponse(barangs)
}

func (service *BarangServiceImpl) Delete(ctx context.Context, barangId int, idUser int) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs := service.BarangRepository.FindById(ctx, tx, barangId, idUser)
	if barangs.Id != barangId {
		panic(exception.NewNotFound("Data Barang Tidak Ditemukan"))
	}
	service.BarangRepository.Delete(ctx, tx, barangs, idUser)
}

func (service *BarangServiceImpl) FindById(ctx context.Context, barangId int, idUser int) web.BarangResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs := service.BarangRepository.FindById(ctx, tx, barangId, idUser)
	if barangs.Id != barangId {
		panic(exception.NewNotFound("Data Barang Tidak Ditemukan"))
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
