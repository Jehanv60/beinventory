package web

type BarangUpdate struct {
	Id         int    `validate:"required"`
	NameProd   string `validate:"required,max=100,min=1" json:"nameprod"`
	Hargaprod  int    `validate:"required,gte=1" json:"hargaprod"`
	Keterangan string `json:"keterangan"`
	Stok       int    `json:"stok"`
}

type PenggunaUpdate struct {
	Id       int    `validate:"required"`
	Pengguna string `validate:"required,max=100,min=1" json:"pengguna"`
	Email    string `validate:"required" json:"email"`
	Sandi    string `validate:"required,max=100,min=1" json:"sandi"`
}
