package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Jehanv60/exception"
	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/domain"
	"github.com/Jehanv60/model/web"
	"github.com/Jehanv60/repository"
	"github.com/Jehanv60/util"
	"github.com/go-playground/validator/v10"
)

type PenggunaServiceImpl struct {
	PenggunaRepository repository.PenggunaRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewPenggunaService(penggunaRepository repository.PenggunaRepository, DB *sql.DB, validate *validator.Validate) PenggunaService {
	return &PenggunaServiceImpl{
		PenggunaRepository: penggunaRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

// Create implements PenggunaService.
func (service *PenggunaServiceImpl) Create(ctx context.Context, request web.PenggunaCreateRequest) web.PenggunaResponse {
	err := service.Validate.Struct(request)
	util.ErrValidateSelf(err)
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	penggunass := service.PenggunaRepository.FindByPenggunaRegister(ctx, tx, request.Pengguna, request.Email)
	if penggunass.Pengguna == request.Pengguna {
		panic(exception.NewSameFound(fmt.Sprintf("%s Sudah Digunakan", request.Pengguna)))
	}
	if penggunass.Email == request.Email {
		panic(exception.NewSameFound(fmt.Sprintf("%s Sudah Digunakan", request.Email)))
	}
	penggunas := domain.Pengguna{
		Pengguna: request.Pengguna,
		Email:    request.Email,
		Sandi:    request.Sandi,
	}
	hashedPass, err := util.Hashpassword(penggunas.Sandi)
	helper.PanicError(err)
	penggunas.Sandi = hashedPass
	penggunas = service.PenggunaRepository.Save(ctx, tx, penggunas)
	return helper.ToPenggunaResponse(penggunas)
}

// FindAll implements PenggunaService.
func (service *PenggunaServiceImpl) FindAll(ctx context.Context) []web.PenggunaResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	penggunas := service.PenggunaRepository.FindAll(ctx, tx)
	return helper.ToPenggunaResponses(penggunas)
}

// FindById implements PenggunaService.
func (service *PenggunaServiceImpl) FindById(ctx context.Context, penggunaId int) web.PenggunaResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	penggunas := service.PenggunaRepository.FindById(ctx, tx, penggunaId)
	if penggunas.Id != penggunaId {
		panic(exception.NewNotFound("Data Tidak Ditemukan"))
	}
	return helper.ToPenggunaResponse(penggunas)
}

// FindById implements PenggunaService.
func (service *PenggunaServiceImpl) FindByPenggunaLogin(ctx context.Context, namaPengguna string) web.PenggunaResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	penggunas := service.PenggunaRepository.FindByPenggunaLogin(ctx, tx, namaPengguna)
	return helper.ToPenggunaResponse(penggunas)
}

// Update implements PenggunaService.
func (service *PenggunaServiceImpl) Update(ctx context.Context, update web.PenggunaUpdate) web.PenggunaResponse {
	service.Validate.RegisterValidation("alphanumdash", util.ValidateAlphanumdash)
	err := service.Validate.Struct(update)
	util.ErrValidateSelf(err)
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	penggunas := service.PenggunaRepository.FindById(ctx, tx, update.Id)
	if penggunas.Id != update.Id {
		panic(exception.NewNotFound("Data Tidak Ditemukan"))
	}
	penggunas.Pengguna = update.Pengguna
	penggunas.Email = update.Email
	penggunas.Sandi = update.Sandi
	hashedPass, err := util.Hashpassword(penggunas.Sandi)
	helper.PanicError(err)
	penggunas.Sandi = hashedPass
	penggunass := service.PenggunaRepository.FindByPenggunaRegister(ctx, tx, update.Pengguna, update.Email)
	if penggunas.Id == penggunass.Id {
		penggunas = service.PenggunaRepository.Update(ctx, tx, penggunas)
		return helper.ToPenggunaResponse(penggunas)
	}
	if penggunas.Id != penggunass.Id {
		if penggunas.Pengguna == penggunass.Pengguna {
			panic(exception.NewSameFound(fmt.Sprintf("%s Sudah Digunakan", update.Pengguna)))
		}
		if penggunas.Email == penggunass.Email {
			panic(exception.NewSameFound(fmt.Sprintf("%s Sudah Digunakan", update.Email)))
		}
	}
	penggunas = service.PenggunaRepository.Update(ctx, tx, penggunas)
	return helper.ToPenggunaResponse(penggunas)
}
