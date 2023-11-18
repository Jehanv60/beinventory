package web

type BarangCreateRequest struct {
	NameProd   string `validate:"required,max=100,min=1" json:"nameprod"`
	Hargaprod  int    `validate:"required" json:"hargaprod"`
	Keterangan string `json:"keterangan"`
	Stok       int    `json:"stok"`
}

type PenggunaCreateRequest struct {
	Pengguna string `validate:"required,max=100,min=1" json:"pengguna"`
	Email    string `validate:"required" json:"email"`
	Sandi    string `validate:"required,max=100,min=1" json:"sandi"`
}
