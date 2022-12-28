package helper

import (
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryCreateResponse {
	return web.CategoryCreateResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryCreateResponse {
	var categoriesResponse []web.CategoryCreateResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, ToCategoryResponse(category))
	}
	return categoriesResponse
}
