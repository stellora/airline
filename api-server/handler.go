package main

import (
	"context"
	"database/sql"

	"github.com/stellora/shop/api-server/api"
	"github.com/stellora/shop/api-server/db"
)

func NewHandler(db *sql.DB, queries *db.Queries) *Handler {
	return &Handler{
		db:      db,
		queries: queries,
	}
}

type Handler struct {
	db      *sql.DB
	queries *db.Queries
}

func (h *Handler) HealthCheck(ctx context.Context, request api.HealthCheckRequestObject) (api.HealthCheckResponseObject, error) {
	ok := true
	return api.HealthCheck200JSONResponse{Ok: &ok}, nil
}

var _ api.StrictServerInterface = (*Handler)(nil)
