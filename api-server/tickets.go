package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func fromDBTicket(t db.GetTicketRow) api.Ticket {
	return api.Ticket{
		Id:        int(t.Ticket.ID),
		Number:    t.Ticket.Number,
		Itinerary: toItinerarySpecs(t.Itinerary),
		Passenger: fromDBPassenger(t.Passenger),
		FareBasis: t.Ticket.FareBasis,
	}
}

func (h *Handler) GetTicket(ctx context.Context, request api.GetTicketRequestObject) (api.GetTicketResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	ticket, err := queriesTx.GetTicket(ctx, db.GetTicketParams{
		Number: sql.NullString{String: request.TicketNumber, Valid: true},
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetTicket404Response{}, nil
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.GetTicket200JSONResponse(fromDBTicket(ticket)), nil
}

func (h *Handler) ListTickets(ctx context.Context, request api.ListTicketsRequestObject) (api.ListTicketsResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	tickets, err := queriesTx.ListTickets(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]api.Ticket, len(tickets))
	for i, t := range tickets {
		result[i] = fromDBTicket(db.GetTicketRow(t))
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.ListTickets200JSONResponse(result), nil
}
