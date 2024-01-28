package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/domain"
)

type BarangRepoImpl struct {
}

func NewRepositoryBarang() BarangRepository {
	return &BarangRepoImpl{}
}

func (repository *BarangRepoImpl) Save(ctx context.Context, tx *sql.Tx, barang domain.Barang) domain.Barang {
	SQL := "insert into barang(nameprod, hargaprod, keterangan, stok)values ($1,$2,$3,$4) returning id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, barang.NameProd, barang.Hargaprod, barang.Keterangan, barang.Stok).Scan(&id)
	helper.PanicError(err)
	barang.Id = id
	return barang
}

func (repository *BarangRepoImpl) Update(ctx context.Context, tx *sql.Tx, barang domain.Barang) domain.Barang {
	SQL := "update barang set nameprod = $2, hargaprod = $3, keterangan = $4, stok = $5 where id = $1 returning id"
	_, err := tx.ExecContext(ctx, SQL, barang.Id, barang.NameProd, barang.Hargaprod, barang.Keterangan, barang.Stok)
	helper.PanicError(err)
	return barang
}

func (repository *BarangRepoImpl) Delete(ctx context.Context, tx *sql.Tx, barang domain.Barang) {
	SQL := "delete from barang where id = $1"
	_, err := tx.ExecContext(ctx, SQL, barang.Id)
	helper.PanicError(err)
}

func (repository *BarangRepoImpl) FindById(ctx context.Context, tx *sql.Tx, barangId int) (domain.Barang, error) {
	SQL := "select id, nameprod, hargaprod, keterangan, stok from barang where id = $1"
	rows, err := tx.QueryContext(ctx, SQL, barangId)
	helper.PanicError(err)
	barang := domain.Barang{}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&barang.Id, &barang.NameProd, &barang.Hargaprod, &barang.Keterangan, &barang.Stok)
		return barang, nil
	} else {
		return barang, errors.New("data barang tidak ditemukan")
	}
}

func (repository *BarangRepoImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Barang {
	SQL := "select id, nameprod, hargaprod, keterangan, stok from barang"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)
	defer rows.Close()
	var barangs []domain.Barang
	for rows.Next() {
		barang := domain.Barang{}
		err := rows.Scan(&barang.Id, &barang.NameProd, &barang.Hargaprod, &barang.Keterangan, &barang.Stok)
		helper.PanicError(err)
		barangs = append(barangs, barang)
	}
	return barangs
}
