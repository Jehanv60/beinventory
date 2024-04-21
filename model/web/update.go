package web

type BarangUpdate struct {
	Id         int    `validate:"required"`
	KodeBarang string `validate:"required,max=100,min=1,alphanum" json:"kodebarang"`
	NameProd   string `validate:"required,max=100,min=1,alphanumdash" json:"nameprod"`
	HargaProd  int    `validate:"required,gte=1" json:"hargaprod"`
	JualProd   int    `validate:"required,gte=1" json:"jualprod"`
	ProfitProd int    `validate:"required,gte=1" json:"profitprod"`
	Keterangan string `validate:"required,max=100,min=1,alphanumdash" json:"keterangan"`
	Stok       int    `validate:"required,gte=1" json:"stok"`
	IdUser     int    `json:"iduser"`
}

type PenggunaUpdate struct {
	Id       int    `validate:"required"`
	Pengguna string `validate:"required,max=100,min=1,alphanumdash" json:"pengguna"`
	Email    string `validate:"required,email" json:"email"`
	Sandi    string `validate:"required,max=100,min=1,alphanumdash" json:"sandi"`
}
