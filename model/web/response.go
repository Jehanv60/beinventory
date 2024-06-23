package web

import "encoding/json"

type BarangResponse struct {
	Id         int    `json:"id"`
	IdUser     int    `json:"iduser"`
	KodeBarang string `json:"kodebarang"`
	NameProd   string `json:"nameprod"`
	HargaProd  int    `json:"HargaProd"`
	JualProd   int    `json:"jualprod"`
	ProfitProd int    `json:"profitprod"`
	Keterangan string `json:"keterangan"`
	Stok       int    `json:"stok"`
}

type PenggunaResponse struct {
	Id       int    `json:"id"`
	Pengguna string `json:"pengguna"`
	Email    string `json:"email"`
	Sandi    string `json:"sandi"`
}

type TransaksiResponse struct {
	Id            int             `json:"id"`
	IdUser        int             `json:"iduser"`
	KodePenjualan string          `json:"kodepenjualan"`
	Jumlah        int             `json:"jumlah"`
	Bayar         int             `json:"bayar"`
	Kembali       int             `json:"kembali"`
	Total         int             `json:"total"`
	Tanggal       string          `json:"tanggal"`
	ItemDetailed  json.RawMessage `json:"itemdetailed"`
}
