package service

import (
	"context"

	"github.com/Jehanv60/model/web"
)

type BarangService interface {
	Create(ctx context.Context, request web.BarangCreateRequest) web.BarangResponse
	Update(ctx context.Context, update web.BarangUpdate) web.BarangResponse
	Delete(ctx context.Context, barangId int)
	FindById(ctx context.Context, barangId int) web.BarangResponse
	FindAll(ctx context.Context) []web.BarangResponse
}
type PenggunaService interface {
	Create(ctx context.Context, request web.PenggunaCreateRequest) web.PenggunaResponse
	Update(ctx context.Context, update web.PenggunaUpdate) web.PenggunaResponse
	FindById(ctx context.Context, penggunaId int) web.PenggunaResponse
	FindAll(ctx context.Context) []web.PenggunaResponse
}
