package repository

import (
	"context"
	"database/sql"

	"github.com/Jehanv60/model/domain"
)

type BarangRepository interface {
	Save(ctx context.Context, tx *sql.Tx, barang domain.Barang, idUser int) domain.Barang
	Update(ctx context.Context, tx *sql.Tx, barang domain.Barang, idUser int) domain.Barang
	Updates(ctx context.Context, tx *sql.Tx, barang []domain.Barang, idUser int) error
	Delete(ctx context.Context, tx *sql.Tx, barang domain.Barang, idUser int)
	FindById(ctx context.Context, tx *sql.Tx, barangId int, idUser int) domain.Barang
	FindByNameRegister(ctx context.Context, tx *sql.Tx, kodeBarang string, barangName string, idUser int) domain.Barang
	FindByNameUpdate(ctx context.Context, tx *sql.Tx, kodeBarang string, barangName string, idUser int) domain.Barang
	FindAll(ctx context.Context, tx *sql.Tx, idUser int) []domain.Barang
}

type PenggunaRepository interface {
	Save(ctx context.Context, tx *sql.Tx, pengguna domain.Pengguna) domain.Pengguna
	Update(ctx context.Context, tx *sql.Tx, pengguna domain.Pengguna) domain.Pengguna
	FindById(ctx context.Context, tx *sql.Tx, penggunaId int) domain.Pengguna
	FindByPenggunaRegister(ctx context.Context, tx *sql.Tx, NamaPengguna, Email string) domain.Pengguna
	FindByPenggunaLogin(ctx context.Context, tx *sql.Tx, NamaPengguna string) domain.Pengguna
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Pengguna
}

type TransaksiRepository interface {
	Save(ctx context.Context, tx *sql.Tx, transaksi domain.Transaction, iduser int) domain.Transaction
	CodeSell(ctx context.Context, tx *sql.Tx, idUser int) []domain.Transaction
	ReportAll(ctx context.Context, tx *sql.Tx, idUser int) []domain.Transaction
}
