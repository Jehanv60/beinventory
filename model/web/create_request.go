package web

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
	IdUser        int    `json:"iduser"`
	Barang        string `validate:"required,max=100,min=1,alphanum" json:"kodebarang"`
	KodePenjualan string `validate:"required,max=100,min=1,alphanum" json:"kodepenjualan"`
	Jumlah        int    `validate:"required,gte=1" json:"jumlah"`
	Bayar         int    `validate:"required,gte=1" json:"bayar"`
	Kembali       int    `validate:"required,gte=1" json:"kembali"`
	Total         int    `validate:"required,gte=1" json:"total"`
	Tanggal       string `json:"tanggal"`
}
