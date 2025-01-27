package web

import "encoding/json"

// alphanumdash berada di util validator.go
type BarangCreateRequest struct {
	KodeBarang string `validate:"required,max=100,min=1,alphanum" json:"kodebarang"`
	JualProd   int    `validate:"required,gte=1" json:"jualprod"`
	NameProd   string `validate:"required,max=100,min=1,alphanumdash" json:"nameprod"`
	HargaProd  int    `validate:"required,gte=1" json:"hargaprod"`
	ProfitProd int    `validate:"required,gte=1" json:"profitprod"`
	Keterangan string `validate:"required,max=100,min=1,alphanumdash" json:"keterangan"`
	Stok       int    `validate:"required,gte=1" json:"stok"`
	IdUser     int    `json:"iduser"`
}

type PenggunaCreateRequest struct {
	Pengguna string `validate:"required,max=100,min=1,alphanum" json:"pengguna"`
	Email    string `validate:"required,email" json:"email"`
	Sandi    string `validate:"required,max=100,min=1,alphanum" json:"sandi"`
}

type TransactionCreateRequest struct {
	Barang json.RawMessage `validate:"required" json:"kodebarang"`
	Bayar  int             `validate:"required,gte=0,lte=1000000000" json:"bayar"`
	// IdUser        int             `json:"iduser"`
	// KodePenjualan string          `validate:"required,max=100,min=1,alphanum" json:"kodepenjualan"`
	// Jumlah        int             `validate:"required,gte=1" json:"jumlah"`
	// Kembali       int             `validate:"required,gte=1" json:"kembali"`
	// Total         int             `validate:"required,gte=1" json:"total"`
	// Tanggal       string          `json:"tanggal"`
	// ItemDetailed  json.RawMessage `json:"itemdetailed"`
}
