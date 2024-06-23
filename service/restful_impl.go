package service

import (
	"context"

	"github.com/Jehanv60/model/web"
)

type BarangService interface {
	Create(ctx context.Context, request web.BarangCreateRequest, idUser int) web.BarangResponse
	Update(ctx context.Context, update web.BarangUpdate, idUser int) web.BarangResponse
	Delete(ctx context.Context, barangId int, idUser int)
	FindById(ctx context.Context, barangId int, idUser int) web.BarangResponse
	FindAll(ctx context.Context, idUser int) []web.BarangResponse
}

type PenggunaService interface {
	Create(ctx context.Context, request web.PenggunaCreateRequest) web.PenggunaResponse
	Update(ctx context.Context, update web.PenggunaUpdate) web.PenggunaResponse
	FindById(ctx context.Context, penggunaId int) web.PenggunaResponse
	FindByPenggunaLogin(ctx context.Context, NamaPengguna string) web.PenggunaResponse
	FindAll(ctx context.Context) []web.PenggunaResponse
}

type TransaksiService interface {
	Create(ctx context.Context, request web.TransactionCreateRequest, idUser int) web.TransaksiResponse
	ReportAll(ctx context.Context, idUser int) []web.TransaksiResponse
}
