-- name: GetAirport :one
SELECT * FROM airports
WHERE id=? LIMIT 1;

-- name: ListAirports :many
SELECT * FROM airports
ORDER BY id ASC;

-- name: CreateAirport :one
INSERT INTO airports (
  iata_code
) VALUES (
  ?
)
RETURNING *;

-- name: UpdateAirport :exec
UPDATE airports SET
iata_code=?
WHERE id=?;

-- name: DeleteAirport :exec
DELETE FROM airports
WHERE id=?;

-- name: DeleteAllAirports :exec
DELETE FROM airports;

-------------------------------------------------------------------------------

-- name: GetFlight :one
SELECT * FROM flights
WHERE id=? LIMIT 1;

-- name: ListFlights :many
SELECT * FROM flights
ORDER BY id ASC;

-- name: CreateFlight :one
INSERT INTO flights (
  number, origin_airport, destination_airport, published
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateFlight :exec
UPDATE flights SET
number=?,
origin_airport=?,
destination_airport=?,
published=?
WHERE id=?;

-- name: DeleteFlight :exec
DELETE FROM flights
WHERE id=?;

-- name: DeleteAllFlights :exec
DELETE FROM flights;