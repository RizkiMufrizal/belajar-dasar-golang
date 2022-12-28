package service

import (
	"belajar-golang-restful-api/model/web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryCreateResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryCreateResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.CategoryCreateResponse
	FindAll(ctx context.Context) []web.CategoryCreateResponse
}
