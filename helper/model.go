package helper

import (
	"github.com/Jehanv60/model/domain"
	"github.com/Jehanv60/model/web"
)

func ToBarangResponse(barang domain.Barang) web.BarangResponse {
	return web.BarangResponse{
		Id:         barang.Id,
		IdUser:     barang.IdUser,
		KodeBarang: barang.KodeBarang,
		NameProd:   barang.NameProd,
		HargaProd:  barang.HargaProd,
		JualProd:   barang.JualProd,
		ProfitProd: barang.ProfitProd,
		Keterangan: barang.Keterangan,
		Stok:       barang.Stok,
	}
}

func ToPenggunaResponse(pengguna domain.Pengguna) web.PenggunaResponse {
	return web.PenggunaResponse{
		Id:       pengguna.Id,
		Pengguna: pengguna.Pengguna,
		Email:    pengguna.Email,
		Sandi:    pengguna.Sandi,
	}
}

func ToBarangResponses(barangs []domain.Barang) []web.BarangResponse {
	var barangResponses []web.BarangResponse
	for _, barangss := range barangs {
		barangResponses = append(barangResponses, ToBarangResponse(barangss))
	}
	return barangResponses
}

func ToPenggunaResponses(penggunas []domain.Pengguna) []web.PenggunaResponse {
	var penggunaResponses []web.PenggunaResponse
	for _, penggunass := range penggunas {
		penggunaResponses = append(penggunaResponses, ToPenggunaResponse(penggunass))
	}
	return penggunaResponses
}

func ToTransaksiResponse(transaksi domain.Transaction) web.TransaksiResponse {
	return web.TransaksiResponse{
		Id:            transaksi.Id,
		IdUser:        transaksi.IdUser,
		KodePenjualan: transaksi.KodePenjualan,
		Jumlah:        transaksi.Jumlah,
		Bayar:         transaksi.Bayar,
		Kembali:       transaksi.Kembali,
		Total:         transaksi.Total,
		Tanggal:       transaksi.Tanggal,
		ItemDetailed:  transaksi.ItemDetailed,
	}
}

func ToTransaksiResponses(transaksis []domain.Transaction) []web.TransaksiResponse {
	var transaksiResponses []web.TransaksiResponse
	for _, transaksiss := range transaksis {
		transaksiResponses = append(transaksiResponses, ToTransaksiResponse(transaksiss))
	}
	return transaksiResponses
}
