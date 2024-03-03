package web

type BarangResponse struct {
	Id         int    `json:"id"`
	IdUser     int    `json:"iduser"`
	NameProd   string `json:"nameprod"`
	Hargaprod  int    `json:"hargaprod"`
	Keterangan string `json:"keterangan"`
	Stok       int    `json:"stok"`
}

type PenggunaResponse struct {
	Id       int    `json:"id"`
	Pengguna string `json:"pengguna"`
	Email    string `json:"email"`
	Sandi    string `json:"sandi"`
}
