------------------------------------------------------------------------------- airports

-- name: GetAirport :one
SELECT * FROM airports
WHERE id=? LIMIT 1;

-- name: GetAirportByIATACode :one
SELECT * FROM airports
WHERE iata_code=? LIMIT 1;

-- name: ListAirports :many
SELECT * FROM airports
ORDER BY id ASC;

-- name: CreateAirport :one
INSERT INTO airports (
  iata_code,
  oadb_id
) VALUES (
  ?, ?
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

------------------------------------------------------------------------------- flights

-- name: GetFlight :one
SELECT * FROM flights_view
WHERE id=? LIMIT 1;

-- name: ListFlights :many
SELECT * FROM flights_view
ORDER BY id ASC;

-- name: CreateFlight :one
INSERT INTO flights (
  number, origin_airport_id, destination_airport_id, published
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateFlight :exec
UPDATE flights SET
number=?,
origin_airport_id=?,
destination_airport_id=?,
published=?
WHERE id=?;

-- name: DeleteFlight :exec
DELETE FROM flights
WHERE id=?;

-- name: DeleteAllFlights :exec
DELETE FROM flights;

------------------------------------------------------------------------------- airport_flights

-- name: ListFlightsByAirport :many
SELECT *
FROM flights_view
WHERE origin_airport_id=:airport OR destination_airport_id=:airport
ORDER BY id ASC;