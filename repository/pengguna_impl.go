package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/domain"
)

type PenggunaRepoImpl struct {
}

func NewRepositoryPengguna() PenggunaRepository {
	return &PenggunaRepoImpl{}
}

func (repository *PenggunaRepoImpl) Save(ctx context.Context, tx *sql.Tx, pengguna domain.Pengguna) domain.Pengguna {
	SQL := "insert into pengguna(pengguna, email, password) values ($1,$2,$3) returning id"
	// result, err := tx.ExecContext(ctx, SQL, barang.NameProd, barang.Hargaprod, barang.Keterangan)
	var id int
	err := tx.QueryRowContext(ctx, SQL, pengguna.Pengguna, pengguna.Email, pengguna.Sandi).Scan(&id)
	helper.PanicError(err)
	pengguna.Id = id
	return pengguna
}

func (repository *PenggunaRepoImpl) Update(ctx context.Context, tx *sql.Tx, pengguna domain.Pengguna) domain.Pengguna {
	SQL := "update pengguna set pengguna = $2, email = $3, password = $4 where id = $1 returning id"
	_, err := tx.ExecContext(ctx, SQL, pengguna.Id, pengguna.Pengguna, pengguna.Email, pengguna.Sandi)
	helper.PanicError(err)
	return pengguna
}

func (repository *PenggunaRepoImpl) FindById(ctx context.Context, tx *sql.Tx, penggunaId int) (domain.Pengguna, error) {
	SQL := "select id, pengguna, email, password from pengguna where id = $1"
	rows, err := tx.QueryContext(ctx, SQL, penggunaId)
	helper.PanicError(err)
	pengguna := domain.Pengguna{}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&pengguna.Id, &pengguna.Pengguna, &pengguna.Email, &pengguna.Sandi)
		return pengguna, nil
	} else {
		return pengguna, errors.New("data User tidak ditemukan")
	}
}

func (repository *PenggunaRepoImpl) FindByPenggunaRegister(ctx context.Context, tx *sql.Tx, NamaPengguna string) (domain.Pengguna, error) {
	SQL := "select id, pengguna, email, password from pengguna where pengguna = $1"
	rows, err := tx.QueryContext(ctx, SQL, NamaPengguna)
	helper.PanicError(err)
	pengguna := domain.Pengguna{}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&pengguna.Id, &pengguna.Pengguna, &pengguna.Email, &pengguna.Sandi)
		if NamaPengguna == pengguna.Pengguna {
			return pengguna, errors.New("pengguna sudah ada")
		}
	}
	return pengguna, nil
}

func (repository *PenggunaRepoImpl) FindByPenggunaLogin(ctx context.Context, tx *sql.Tx, NamaPengguna string) (domain.Pengguna, error) {
	SQL := "select id, pengguna, email, password from pengguna where pengguna = $1"
	rows, err := tx.QueryContext(ctx, SQL, NamaPengguna)
	helper.PanicError(err)
	pengguna := domain.Pengguna{}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&pengguna.Id, &pengguna.Pengguna, &pengguna.Email, &pengguna.Sandi)
	}
	return pengguna, nil
}

func (repository *PenggunaRepoImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Pengguna {
	SQL := "select id, pengguna, email, password from pengguna"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)
	defer rows.Close()
	var penggunas []domain.Pengguna
	for rows.Next() {
		pengguna := domain.Pengguna{}
		err := rows.Scan(&pengguna.Id, &pengguna.Pengguna, &pengguna.Email, &pengguna.Sandi)
		helper.PanicError(err)
		penggunas = append(penggunas, pengguna)
	}
	return penggunas
}
