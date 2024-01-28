package helper

import (
	"github.com/Jehanv60/model/domain"
	"github.com/Jehanv60/model/web"
)

func ToBarangResponse(barang domain.Barang) web.BarangResponse {
	return web.BarangResponse{
		Id:         barang.Id,
		NameProd:   barang.NameProd,
		Hargaprod:  barang.Hargaprod,
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
