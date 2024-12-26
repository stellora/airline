package main

import (
	"context"

	"github.com/stellora/shop/api-server/api"
)

func NewHandler() *Handler {
	return &Handler{}
}

type Handler struct{}

func (h *Handler) GetCategories(ctx context.Context, request api.GetCategoriesRequestObject) (api.GetCategoriesResponseObject, error) {
	return api.GetCategories200JSONResponse{}, nil
}

func (h *Handler) PostCategories(ctx context.Context, request api.PostCategoriesRequestObject) (api.PostCategoriesResponseObject, error) {
	return api.PostCategories201Response{}, nil
}

func (h *Handler) GetCategoriesCategoryIdProducts(ctx context.Context, request api.GetCategoriesCategoryIdProductsRequestObject) (api.GetCategoriesCategoryIdProductsResponseObject, error) {
	return api.GetCategoriesCategoryIdProducts200JSONResponse{}, nil
}

func (h *Handler) DeleteCategoriesId(ctx context.Context, request api.DeleteCategoriesIdRequestObject) (api.DeleteCategoriesIdResponseObject, error) {
	return api.DeleteCategoriesId204Response{}, nil
}

func (h *Handler) GetHealth(ctx context.Context, request api.GetHealthRequestObject) (api.GetHealthResponseObject, error) {
	ok := true
	return api.GetHealth200JSONResponse{Ok: &ok}, nil
}

func (h *Handler) DeleteProducts(ctx context.Context, request api.DeleteProductsRequestObject) (api.DeleteProductsResponseObject, error) {
	return api.DeleteProducts204Response{}, nil
}

func (h *Handler) GetProducts(ctx context.Context, request api.GetProductsRequestObject) (api.GetProductsResponseObject, error) {
	return api.GetProducts200JSONResponse{}, nil
}

func (h *Handler) PostProducts(ctx context.Context, request api.PostProductsRequestObject) (api.PostProductsResponseObject, error) {
	return api.PostProducts201Response{}, nil
}

func (h *Handler) DeleteProductsId(ctx context.Context, request api.DeleteProductsIdRequestObject) (api.DeleteProductsIdResponseObject, error) {
	return api.DeleteProductsId204Response{}, nil
}

func (h *Handler) PatchProductsId(ctx context.Context, request api.PatchProductsIdRequestObject) (api.PatchProductsIdResponseObject, error) {
	return api.PatchProductsId200Response{}, nil
}

func (h *Handler) PutProductsProductIdCategoriesCategoryId(ctx context.Context, request api.PutProductsProductIdCategoriesCategoryIdRequestObject) (api.PutProductsProductIdCategoriesCategoryIdResponseObject, error) {
	return api.PutProductsProductIdCategoriesCategoryId200Response{}, nil
}

var _ api.StrictServerInterface = (*Handler)(nil)
