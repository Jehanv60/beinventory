package repository

import (
	"context"
	"database/sql"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/domain"
)

type TransaksiRepoImpl struct {
}

func NewRepositoryTransaksi() TransaksiRepository {
	return &TransaksiRepoImpl{}
}
func (repository *TransaksiRepoImpl) Save(ctx context.Context, tx *sql.Tx, transaksi domain.Transaction, idUser int) domain.Transaction {
	SQL := "insert into transaksi(kodebarang, iduser, kodepenjualan, jumlah, bayar, kembali, total, tanggal)values ($1,$2,$3,$4,$5,$6,$7,$8) returning id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, transaksi.Barang.KodeBarang, idUser, transaksi.KodePenjualan, transaksi.Jumlah, transaksi.Bayar, transaksi.Kembali, transaksi.Total, transaksi.Tanggal).Scan(&id)
	helper.PanicError(err)
	transaksi.Id = id
	transaksi.IdUser = idUser
	return transaksi
}
