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

func (repository *TransaksiRepoImpl) CodeSell(ctx context.Context, tx *sql.Tx, idUser int) []domain.Transaction {
	SQL := "select id, kodebarang, iduser, kodepenjualan, jumlah, bayar, kembali, total, tanggal from transaksi where iduser=$1 and DATE_PART('month', tanggal) = DATE_PART('month', CURRENT_DATE)"
	rows, err := tx.QueryContext(ctx, SQL, idUser)
	helper.PanicError(err)
	defer rows.Close()
	var transaksiall []domain.Transaction
	for rows.Next() {
		transaksi := domain.Transaction{}
		err := rows.Scan(&transaksi.Id, &transaksi.Barang.KodeBarang, &transaksi.IdUser, &transaksi.KodePenjualan, &transaksi.Jumlah, &transaksi.Bayar, &transaksi.Kembali, &transaksi.Total, &transaksi.Tanggal)
		helper.PanicError(err)
		transaksiall = append(transaksiall, transaksi)
	}
	return transaksiall
}
