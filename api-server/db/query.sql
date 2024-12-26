-- name: GetAirport :one
SELECT * FROM airports
WHERE id=? LIMIT 1;

-- name: ListAirports :many
SELECT * FROM airports
ORDER BY id ASC;

-- name: CreateAirport :one
INSERT INTO airports (
  title
) VALUES (
  ?
)
RETURNING *;

-- name: UpdateAirport :exec
UPDATE airports SET
title=?
WHERE id=?;

-- name: DeleteAirport :exec
DELETE FROM airports
WHERE id=?;

-------------------------------------------------------------------------------

-- name: GetFlight :one
SELECT * FROM flights
WHERE id=? LIMIT 1;

-- name: ListFlights :many
SELECT * FROM flights
ORDER BY id ASC;

-- name: CreateFlight :one
INSERT INTO flights (
  title, published
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateFlight :exec
UPDATE flights SET
title=?,
published=?
WHERE id=?;

-- name: DeleteFlight :exec
DELETE FROM flights
WHERE id=?;

-------------------------------------------------------------------------------

-- name: AddFlightToAirport :exec
INSERT INTO flight_airports (
  flight_id,
  airport_id
) VALUES (
  ?, ?
);

-- name: RemoveFlightFromAirport :exec
DELETE FROM flight_airports
WHERE flight_id = ? AND airport_id = ?;
