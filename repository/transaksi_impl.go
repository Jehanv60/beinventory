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
	SQL := "insert into transaksi(iduser, kodepenjualan, jumlah, bayar, kembali, total, tanggal, itemdetailed)values ($1,$2,$3,$4,$5,$6,$7,$8) returning id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, idUser, transaksi.KodePenjualan, transaksi.Jumlah, transaksi.Bayar, transaksi.Kembali, transaksi.Total, transaksi.Tanggal, transaksi.ItemDetailed).Scan(&id)
	helper.PanicError(err)
	transaksi.Id = id
	transaksi.IdUser = idUser
	return transaksi
}

func (repository *TransaksiRepoImpl) CodeSell(ctx context.Context, tx *sql.Tx, idUser int) []domain.Transaction {
	SQL := "select id, iduser, tanggal from transaksi where iduser=$1 and DATE_PART('month', tanggal) = DATE_PART('month', CURRENT_DATE)"
	rows, err := tx.QueryContext(ctx, SQL, idUser)
	helper.PanicError(err)
	defer rows.Close()
	var codeSell []domain.Transaction
	for rows.Next() {
		transaksi := domain.Transaction{}
		err := rows.Scan(&transaksi.Id, &transaksi.IdUser, &transaksi.Tanggal)
		helper.PanicError(err)
		codeSell = append(codeSell, transaksi)
	}
	return codeSell
}

func (repository *TransaksiRepoImpl) ReportAll(ctx context.Context, tx *sql.Tx, idUser int) []domain.Transaction {
	SQL := `select id, iduser, kodepenjualan, jumlah, bayar, kembali, total, tanggal, itemdetailed from transaksi where iduser = $1 order by id asc;`
	rows, err := tx.QueryContext(ctx, SQL, idUser)
	helper.PanicError(err)
	defer rows.Close()
	var transaksiAll []domain.Transaction
	for rows.Next() {
		transaksi := domain.Transaction{}
		err := rows.Scan(
			&transaksi.Id,
			&transaksi.IdUser,
			&transaksi.KodePenjualan,
			&transaksi.Jumlah,
			&transaksi.Bayar,
			&transaksi.Kembali,
			&transaksi.Total,
			&transaksi.Tanggal,
			&transaksi.ItemDetailed,
		)
		helper.PanicError(err)
		transaksiAll = append(transaksiAll, transaksi)
	}
	return transaksiAll
}
