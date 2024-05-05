package domain

type Transaction struct {
	Id            int
	IdUser        int
	Barang        Barang
	KodePenjualan string
	Jumlah        int
	Bayar         int
	Kembali       int
	Total         int
	Tanggal       string
}
