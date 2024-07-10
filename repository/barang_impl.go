package repository

import (
	"context"
	"database/sql"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/domain"
)

type BarangRepoImpl struct {
}

func NewRepositoryBarang() BarangRepository {
	return &BarangRepoImpl{}
}

func (repository *BarangRepoImpl) Save(ctx context.Context, tx *sql.Tx, barang domain.Barang, idUser int) domain.Barang {
	SQL := "insert into barang(nameprod, HargaProd, keterangan, stok, iduser, kodebarang, jualprod, profitprod)values ($1,$2,$3,$4,$5,$6,$7,$8) returning id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, barang.NameProd, barang.HargaProd, barang.Keterangan, barang.Stok, idUser, barang.KodeBarang, barang.JualProd, barang.ProfitProd).Scan(&id)
	helper.PanicError(err)
	barang.Id = id
	barang.IdUser = idUser
	return barang
}

func (repository *BarangRepoImpl) Update(ctx context.Context, tx *sql.Tx, barang domain.Barang, idUser int) domain.Barang {
	SQL := "update barang set nameprod = $2, HargaProd = $3, keterangan = $4, stok = $5, kodebarang = $6, jualprod = $7, profitprod = $8 where id = $1 and iduser = $9 returning id"
	_, err := tx.ExecContext(ctx, SQL, barang.Id, barang.NameProd, barang.HargaProd, barang.Keterangan, barang.Stok, barang.KodeBarang, barang.JualProd, barang.ProfitProd, idUser)
	helper.PanicError(err)
	return barang
}

func (repository *BarangRepoImpl) Updates(ctx context.Context, tx *sql.Tx, barang []domain.Barang, idUser int) error {
	SQL := "update barang set nameprod = $2, HargaProd = $3, keterangan = $4, stok = $5, kodebarang = $6, jualprod = $7, profitprod = $8 where id = $1 and iduser = $9 returning id"
	hasil, err := tx.PrepareContext(ctx, SQL)
	helper.PanicError(err)
	defer hasil.Close()
	for _, v := range barang {
		_, err := hasil.ExecContext(ctx, v.Id, v.NameProd, v.HargaProd, v.Keterangan, v.Stok, v.KodeBarang, v.JualProd, v.ProfitProd, idUser)
		helper.PanicError(err)
	}
	return nil
}

func (repository *BarangRepoImpl) Delete(ctx context.Context, tx *sql.Tx, barang domain.Barang, idUser int) {
	SQL := "delete from barang where id = $1 and iduser=$2"
	_, err := tx.ExecContext(ctx, SQL, barang.Id, idUser)
	helper.PanicError(err)
}

func (repository *BarangRepoImpl) FindById(ctx context.Context, tx *sql.Tx, barangId int, idUser int) domain.Barang {
	SQL := "select id, iduser, nameprod, HargaProd, keterangan, stok, kodebarang, jualprod, profitprod from barang where id = $1 and iduser=$2"
	rows, err := tx.QueryContext(ctx, SQL, barangId, idUser)
	helper.PanicError(err)
	barang := domain.Barang{}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&barang.Id, &barang.IdUser, &barang.NameProd, &barang.HargaProd, &barang.Keterangan, &barang.Stok, &barang.KodeBarang, &barang.JualProd, &barang.ProfitProd)
	}
	return barang
}

func (repository *BarangRepoImpl) FindByNameRegister(ctx context.Context, tx *sql.Tx, kodeBarang string, barangName string, idUser int) domain.Barang {
	SQL := "select id, iduser, nameprod, HargaProd, keterangan, stok, kodebarang, jualprod, profitprod from barang where nameprod = $1 and iduser = $3 or kodebarang = $2 and iduser = $3"
	rows, err := tx.QueryContext(ctx, SQL, barangName, kodeBarang, idUser)
	helper.PanicError(err)
	barang := domain.Barang{}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&barang.Id, &barang.IdUser, &barang.NameProd, &barang.HargaProd, &barang.Keterangan, &barang.Stok, &barang.KodeBarang, &barang.JualProd, &barang.ProfitProd)
	}
	return barang
}

func (repository *BarangRepoImpl) FindByNameUpdate(ctx context.Context, tx *sql.Tx, kodeBarang string, barangName string, idUser int) domain.Barang {
	SQL := "select id, iduser, nameprod, HargaProd, keterangan, stok, kodebarang, jualprod, profitprod from barang where nameprod = $1 and iduser = $3 or kodebarang = $2 and iduser = $3"
	rows, err := tx.QueryContext(ctx, SQL, barangName, kodeBarang, idUser)
	helper.PanicError(err)
	barang := domain.Barang{}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&barang.Id, &barang.IdUser, &barang.NameProd, &barang.HargaProd, &barang.Keterangan, &barang.Stok, &barang.KodeBarang, &barang.JualProd, &barang.ProfitProd)
	}
	return barang
}

func (repository *BarangRepoImpl) FindAll(ctx context.Context, tx *sql.Tx, idUser int) []domain.Barang {
	SQL := "select id, iduser, nameprod, HargaProd, keterangan, stok, kodebarang, jualprod, profitprod from barang where iduser=$1"
	rows, err := tx.QueryContext(ctx, SQL, idUser)
	helper.PanicError(err)
	defer rows.Close()
	var barangs []domain.Barang
	for rows.Next() {
		barang := domain.Barang{}
		err := rows.Scan(&barang.Id, &barang.IdUser, &barang.NameProd, &barang.HargaProd, &barang.Keterangan, &barang.Stok, &barang.KodeBarang, &barang.JualProd, &barang.ProfitProd)
		helper.PanicError(err)
		barangs = append(barangs, barang)
	}
	return barangs
}
