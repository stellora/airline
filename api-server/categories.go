package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/stellora/shop/api-server/api"
)

var (
	categories = []api.Category{
		{Title: "Silverware"},
		{Title: "Cookware"},
		{Title: "Vegetables"},
	}
)

func getCategory(id string) *api.Category {
	for i := range categories {
		if categories[i].Id == id {
			return &categories[i]
		}
	}
	return nil
}

func init() {
	for i := range categories {
		categories[i].Id = strconv.Itoa(i + 1)
	}
}

func (h *Handler) GetCategory(ctx context.Context, request api.GetCategoryRequestObject) (api.GetCategoryResponseObject, error) {
	category := getCategory(request.Id)
	if category == nil {
		return &api.GetCategory404Response{}, nil
	}
	return api.GetCategory200JSONResponse(*category), nil
}

func (h *Handler) ListCategories(ctx context.Context, request api.ListCategoriesRequestObject) (api.ListCategoriesResponseObject, error) {
	return api.ListCategories200JSONResponse(categories), nil
}

func (h *Handler) CreateCategory(ctx context.Context, request api.CreateCategoryRequestObject) (api.CreateCategoryResponseObject, error) {
	title := request.Body.Title
	if title == "" {
		return nil, fmt.Errorf("title must not be empty")
	}

	for _, category := range categories {
		if category.Title == title {
			return nil, fmt.Errorf("title must be unique across all categories")
		}
	}

	newCategory := api.Category{
		Id:    strconv.Itoa(len(categories) + 1),
		Title: title,
	}
	categories = append(categories, newCategory)

	return api.CreateCategory201Response{}, nil
}

func (h *Handler) DeleteCategory(ctx context.Context, request api.DeleteCategoryRequestObject) (api.DeleteCategoryResponseObject, error) {
	// Find and remove the category
	newCategories := []api.Category{}
	found := false
	for _, category := range categories {
		if category.Id != request.Id {
			newCategories = append(newCategories, category)
		} else {
			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf("category not found")
	}

	// Remove all product memberships for this category
	newMemberships := []productCategoryMembership{}
	for _, membership := range productCategoryMemberships {
		if membership.category != request.Id {
			newMemberships = append(newMemberships, membership)
		}
	}

	categories = newCategories
	productCategoryMemberships = newMemberships

	return api.DeleteCategory204Response{}, nil
}

var _ api.StrictServerInterface = (*Handler)(nil)
