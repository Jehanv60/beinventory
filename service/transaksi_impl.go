package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Jehanv60/exception"
	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/domain"
	"github.com/Jehanv60/model/web"
	"github.com/Jehanv60/repository"
	"github.com/Jehanv60/util"
	"github.com/go-playground/validator/v10"
)

type TransaksiServiceImpl struct {
	TransaksiRepository repository.TransaksiRepository
	BarangRepository    repository.BarangRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewTransaksiService(transaksiRepository repository.TransaksiRepository, barangRepository repository.BarangRepository, DB *sql.DB, validate *validator.Validate) TransaksiService {
	return &TransaksiServiceImpl{
		TransaksiRepository: transaksiRepository,
		BarangRepository:    barangRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *TransaksiServiceImpl) Create(ctx context.Context, request web.TransactionCreateRequest, idUser int) web.TransaksiResponse {
	service.Validate.RegisterValidation("alphanumdash", util.ValidateAlphanumdash)
	err := service.Validate.Struct(request)
	util.ErrValidateSelf(err)
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	barangs, _ := service.BarangRepository.FindByNameRegister(ctx, tx, request.Barang, request.Barang, idUser)
	zone, _ := time.LoadLocation("Asia/Jakarta")
	transaksi := domain.Transaction{
		IdUser: idUser,
		Barang: domain.Barang{
			KodeBarang: request.Barang,
		},
		KodePenjualan: request.KodePenjualan,
		Jumlah:        request.Jumlah,
		Bayar:         request.Bayar,
		Kembali:       request.Bayar - request.Total,
		Total:         barangs.JualProd * request.Jumlah,
		Tanggal:       time.Now().UTC().In(zone).Format(("2006-01-02 15:04:05")),
	}
	if barangs.KodeBarang != transaksi.Barang.KodeBarang {
		panic(exception.NewNotFound(fmt.Sprintf("%s Data Barang Tidak Ada, Mohon Untuk Cek Di Inventory", request.Barang)))
	}
	if transaksi.Bayar < transaksi.Total {
		panic(exception.NewNotEqual("Uang Gak Cukup"))
	}
	if transaksi.Jumlah > barangs.Stok {
		panic(exception.NewNotEqual("Stok Gak Cukup"))
	}
	transaksi.Kembali = transaksi.Bayar - transaksi.Total
	barangs.Stok = barangs.Stok - request.Jumlah
	service.BarangRepository.Update(ctx, tx, barangs, idUser)
	transaksi = service.TransaksiRepository.Save(ctx, tx, transaksi, idUser)
	return helper.ToTransaksiResponse(transaksi)
}
