package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/domain"
)

type BarangRepoImpl struct {
}

func NewRepositoryBarang() BarangRepository {
	return &BarangRepoImpl{}
}

func (repository *BarangRepoImpl) Save(ctx context.Context, tx *sql.Tx, barang domain.Barang, idUser int) domain.Barang {
	SQL := "insert into barang(nameprod, hargaprod, keterangan, stok, iduser)values ($1,$2,$3,$4,$5) returning id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, barang.NameProd, barang.Hargaprod, barang.Keterangan, barang.Stok, idUser).Scan(&id)
	helper.PanicError(err)
	barang.Id = id
	barang.IdUser = idUser
	return barang
}

func (repository *BarangRepoImpl) Update(ctx context.Context, tx *sql.Tx, barang domain.Barang, idUser int) domain.Barang {
	SQL := "update barang set nameprod = $2, hargaprod = $3, keterangan = $4, stok = $5 where id = $1 and iduser = $6 returning id"
	_, err := tx.ExecContext(ctx, SQL, barang.Id, barang.NameProd, barang.Hargaprod, barang.Keterangan, barang.Stok, idUser)
	helper.PanicError(err)
	return barang
}

func (repository *BarangRepoImpl) Delete(ctx context.Context, tx *sql.Tx, barang domain.Barang, idUser int) {
	SQL := "delete from barang where id = $1 and iduser=$2"
	_, err := tx.ExecContext(ctx, SQL, barang.Id, idUser)
	helper.PanicError(err)
}

func (repository *BarangRepoImpl) FindById(ctx context.Context, tx *sql.Tx, barangId int, idUser int) (domain.Barang, error) {
	SQL := "select id, iduser, nameprod, hargaprod, keterangan, stok from barang where id = $1 and iduser=$2"
	rows, err := tx.QueryContext(ctx, SQL, barangId, idUser)
	helper.PanicError(err)
	barang := domain.Barang{}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&barang.Id, &barang.IdUser, &barang.NameProd, &barang.Hargaprod, &barang.Keterangan, &barang.Stok)
		return barang, nil
	} else {
		return barang, errors.New("data barang tidak ditemukan")
	}
}

func (repository *BarangRepoImpl) FindByNameRegister(ctx context.Context, tx *sql.Tx, barangName string, idUser int) (domain.Barang, error) {
	SQL := "select id, iduser, nameprod, hargaprod, keterangan, stok from barang where nameprod = $1 and iduser=$2"
	rows, err := tx.QueryContext(ctx, SQL, barangName, idUser)
	helper.PanicError(err)
	namas := "Nama Barang"
	barang := domain.Barang{}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&barang.Id, &barang.IdUser, &barang.NameProd, &barang.Hargaprod, &barang.Keterangan, &barang.Stok)
		if barangName == barang.NameProd {
			return barang, fmt.Errorf("%s %s Sudah Digunakan, Mohon Untuk Cek Di Inventory", namas, barangName)
		}
	}
	return barang, nil
}

func (repository *BarangRepoImpl) FindByNameUpdate(ctx context.Context, tx *sql.Tx, barangName string, idUser int) (domain.Barang, error) {
	SQL := "select id, iduser, nameprod, hargaprod, keterangan, stok from barang where nameprod = $1 and iduser=$2"
	rows, err := tx.QueryContext(ctx, SQL, barangName, idUser)
	helper.PanicError(err)
	namas := "Nama Barang"
	barang := domain.Barang{}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&barang.Id, &barang.IdUser, &barang.NameProd, &barang.Hargaprod, &barang.Keterangan, &barang.Stok)
		if barangName == barang.NameProd {
			return barang, fmt.Errorf("%s %s Sudah Digunakan, Mohon Diganti Dengan Yang lain", namas, barangName)
		}
	}
	return barang, nil

}

func (repository *BarangRepoImpl) FindAll(ctx context.Context, tx *sql.Tx, idUser int) []domain.Barang {
	SQL := "select id, iduser, nameprod, hargaprod, keterangan, stok from barang where iduser=$1"
	rows, err := tx.QueryContext(ctx, SQL, idUser)
	helper.PanicError(err)
	defer rows.Close()
	var barangs []domain.Barang
	for rows.Next() {
		barang := domain.Barang{}
		err := rows.Scan(&barang.Id, &barang.IdUser, &barang.NameProd, &barang.Hargaprod, &barang.Keterangan, &barang.Stok)
		helper.PanicError(err)
		barangs = append(barangs, barang)
	}
	return barangs
}
