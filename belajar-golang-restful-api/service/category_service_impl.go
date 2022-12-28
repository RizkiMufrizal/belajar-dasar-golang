package service

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (categoryService *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryCreateResponse {
	err := categoryService.Validate.Struct(request)
	helper.PanifIfError(err)

	tx, err := categoryService.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}
	category = categoryService.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (categoryService *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryCreateResponse {
	err := categoryService.Validate.Struct(request)
	helper.PanifIfError(err)

	tx, err := categoryService.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := categoryService.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanifIfError(err)

	category.Name = request.Name
	category = categoryService.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (categoryService *CategoryServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := categoryService.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := categoryService.CategoryRepository.FindById(ctx, tx, id)
	helper.PanifIfError(err)

	categoryService.CategoryRepository.Delete(ctx, tx, category)
}

func (categoryService *CategoryServiceImpl) FindById(ctx context.Context, id int) web.CategoryCreateResponse {
	tx, err := categoryService.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := categoryService.CategoryRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(err)
	}

	return helper.ToCategoryResponse(category)
}

func (categoryService *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryCreateResponse {
	tx, err := categoryService.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := categoryService.CategoryRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categories)
}
