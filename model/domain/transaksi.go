package domain

import "encoding/json"

type Transaction struct {
	Id            int
	IdUser        int
	KodePenjualan string
	Jumlah        int
	Bayar         int
	Kembali       int
	Total         int
	Tanggal       string
	ItemDetailed  json.RawMessage
}

type Product struct {
	KodeProd string `json:"kodeprod"`
	Jumlah   int    `json:"jumlah"`
}
