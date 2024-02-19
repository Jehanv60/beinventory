package web

type BarangUpdate struct {
	Id         int    `validate:"required"`
	NameProd   string `validate:"required,max=100,min=1,alphanumdash" json:"nameprod"`
	Hargaprod  int    `validate:"required,gte=1" json:"hargaprod"`
	Keterangan string `validate:"required,max=100,min=1,alphanumdash" json:"keterangan"`
	Stok       int    `validate:"required,gte=1" json:"stok"`
}

type PenggunaUpdate struct {
	Id       int    `validate:"required"`
	Pengguna string `validate:"required,max=100,min=1,alphanumdash" json:"pengguna"`
	Email    string `validate:"required,email" json:"email"`
	Sandi    string `validate:"required,max=100,min=1,alphanumdash" json:"sandi"`
}
