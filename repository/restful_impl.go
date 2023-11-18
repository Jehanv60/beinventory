package repository

import (
	"context"
	"database/sql"

	"github.com/Jehanv60/model/domain"
)

type BarangRepository interface {
	Save(ctx context.Context, tx *sql.Tx, barang domain.Barang) domain.Barang
	Update(ctx context.Context, tx *sql.Tx, barang domain.Barang) domain.Barang
	Delete(ctx context.Context, tx *sql.Tx, barang domain.Barang)
	FindById(ctx context.Context, tx *sql.Tx, barangId int) (domain.Barang, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Barang
}

type PenggunaRepository interface {
	Save(ctx context.Context, tx *sql.Tx, pengguna domain.Pengguna) domain.Pengguna
	Update(ctx context.Context, tx *sql.Tx, pengguna domain.Pengguna) domain.Pengguna
	FindById(ctx context.Context, tx *sql.Tx, penggunaId int) (domain.Pengguna, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Pengguna
}
