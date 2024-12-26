package main

import (
	"context"

	"github.com/stellora/shop/api-server/api"
)

func NewHandler() *Handler {
	return &Handler{}
}

type Handler struct{}

func (h *Handler) HealthCheck(ctx context.Context, request api.HealthCheckRequestObject) (api.HealthCheckResponseObject, error) {
	ok := true
	return api.HealthCheck200JSONResponse{Ok: &ok}, nil
}

var _ api.StrictServerInterface = (*Handler)(nil)
