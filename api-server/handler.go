package main

import (
	"context"

	"github.com/stellora/shop/api-server/api"
)

func NewHandler() *Handler {
	return &Handler{}
}

type Handler struct{}

func (h *Handler) GetHealth(ctx context.Context, request api.GetHealthRequestObject) (api.GetHealthResponseObject, error) {
	ok := true
	return api.GetHealth200JSONResponse{Ok: &ok}, nil
}

var _ api.StrictServerInterface = (*Handler)(nil)
