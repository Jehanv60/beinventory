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
		zone, _    = time.LoadLocation("Asia/Jakarta")
		produk     []domain.Product
		barangb    []domain.Barang
		transaksi  = domain.Transaction{}
		allDetail  []map[string]interface{}
		rawMessage json.RawMessage
		sum        int
		total      int
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
		if v.Jumlah > barangs.Stok {
			panic(exception.NewNotEqual(fmt.Sprintf("%s Stok Tidak Cukup", v.KodeProd)))
		}
		if v.Jumlah <= 0 {
			panic(exception.NewNotEqual(fmt.Sprintf("%s Jumlah Barang Harus Lebih Dari 0", v.KodeProd)))
		}
		detail := map[string]interface{}{
			"id":       barangs.Id,
			"jumlah":   v.Jumlah,
			"kodeprod": barangs.KodeBarang,
		}
		barangb = append(barangb, barangs)
		allDetail = append(allDetail, detail)
	}
	updated := service.BarangRepository.Updates(ctx, tx, barangb, idUser)
	helper.PanicError(updated)
	data, err := json.Marshal(allDetail)
	helper.PanicError(err)
	json.Unmarshal([]byte(rawMessage), &allDetail)
	rawMessage = data
	transaksi = domain.Transaction{
		IdUser:  idUser,
		Jumlah:  sum,
		Bayar:   request.Bayar,
		Tanggal: time.Now().UTC().In(zone).Format(("2006-01-02 15:04:05")),
	}
	transaksi.ItemDetailed = rawMessage
	transaksi.Total = total
	if transaksi.Bayar < transaksi.Total {
		panic(exception.NewNotEqual(fmt.Sprintf("Uang Gak Cukup kurang Rp %v", transaksi.Total-transaksi.Bayar)))
	}
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
