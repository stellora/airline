package main

import (
	"context"
	"fmt"

	"github.com/stellora/airline/api-server/api"
)

var (
	productCategoryMemberships = []productCategoryMembership{
		{product: "1", category: "1"},
		{product: "2", category: "1"},
		{product: "3", category: "1"},
		{product: "4", category: "2"},
		{product: "5", category: "2"},
		{product: "6", category: "2"},
		{product: "7", category: "3"},
		{product: "8", category: "3"},
		{product: "9", category: "3"},
	}
)

type productCategoryMembership struct {
	product, category string
}

func (h *Handler) UpdateProductCategoryMembership(ctx context.Context, request api.UpdateProductCategoryMembershipRequestObject) (api.UpdateProductCategoryMembershipResponseObject, error) {
	// Find product
	var productFound bool
	for _, p := range products {
		if p.Id == request.ProductId {
			productFound = true
			break
		}
	}
	if !productFound {
		return nil, fmt.Errorf("product with id %q not found", request.ProductId)
	}

	// Find category
	var categoryFound bool
	for _, c := range categories {
		if c.Id == request.CategoryId {
			categoryFound = true
			break
		}
	}
	if !categoryFound {
		return nil, fmt.Errorf("category with id %q not found", request.CategoryId)
	}

	// Find existing membership
	existingIndex := -1
	for i, m := range productCategoryMemberships {
		if m.product == request.ProductId && m.category == request.CategoryId {
			existingIndex = i
			break
		}
	}

	if request.Body.Value {
		if existingIndex == -1 {
			productCategoryMemberships = append(productCategoryMemberships, productCategoryMembership{
				product:  request.ProductId,
				category: request.CategoryId,
			})
		}
	} else {
		if existingIndex != -1 {
			productCategoryMemberships = append(productCategoryMemberships[:existingIndex], productCategoryMemberships[existingIndex+1:]...)
		}
	}

	return api.UpdateProductCategoryMembership204Response{}, nil
}
