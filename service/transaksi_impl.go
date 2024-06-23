package service

import (
	"context"
	"database/sql"
	"encoding/json"
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
	// service.Validate.RegisterValidation("alphanumdash", util.ValidateAlphanumdash)
	// err := service.Validate.Struct(request)
	// util.ErrValidateSelf(err)
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	var (
		zone, _   = time.LoadLocation("Asia/Jakarta")
		produk    []domain.Product
		result    []domain.Barang
		transaksi = domain.Transaction{}
		sum       int
		total     int
	)
	json.Unmarshal([]byte(request.Barang), &produk)
	for _, v := range produk {
		barangs := service.BarangRepository.FindByNameRegister(ctx, tx, v.KodeProd, "", idUser)
		if barangs.KodeBarang != v.KodeProd {
			panic(exception.NewNotFound(fmt.Sprintf("%s Data Barang Tidak Ada, Mohon Untuk Cek Di Inventory", v.KodeProd)))
		}
		sum += v.Jumlah
		total += barangs.JualProd * v.Jumlah
		barangs.Stok = barangs.Stok - v.Jumlah
		if transaksi.Jumlah > barangs.Stok {
			panic(exception.NewNotEqual(fmt.Sprintf("%s Stok Tidak Cukup", v.KodeProd)))
		}
		result = append(result, barangs)
	}
	for _, v := range result {
		transaksi = domain.Transaction{
			IdUser:  idUser,
			Jumlah:  sum,
			Bayar:   request.Bayar,
			Tanggal: time.Now().UTC().In(zone).Format(("2006-01-02 15:04:05")),
		}
		service.BarangRepository.Update(ctx, tx, v, idUser)
	}
	transaksi.Total = total
	if transaksi.Bayar < transaksi.Total {
		panic(exception.NewNotEqual(fmt.Sprintf("Uang Gak Cukup kurang Rp %v", transaksi.Total-transaksi.Bayar)))
	}
	transaksi.ItemDetailed = request.Barang
	transaksi.Kembali = transaksi.Bayar - transaksi.Total
	countId := service.TransaksiRepository.CodeSell(ctx, tx, idUser)
	transaksi.KodePenjualan = fmt.Sprintf("PJ/%v/%s", util.ChangeMonth(countId), time.Now().UTC().In(zone).Format(("06-01-02")))
	transaksi = service.TransaksiRepository.Save(ctx, tx, transaksi, idUser)
	return helper.ToTransaksiResponse(transaksi)
}
func (service *TransaksiServiceImpl) ReportAll(ctx context.Context, idUser int) []web.TransaksiResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	transaksi := service.TransaksiRepository.ReportAll(ctx, tx, idUser)
	return helper.ToTransaksiResponses(transaksi)
}
