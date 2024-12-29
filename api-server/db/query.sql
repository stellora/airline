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

-- name: UpdateAirport :one
UPDATE airports SET
iata_code = COALESCE(sqlc.narg('iata_code'), iata_code)
WHERE id=?
RETURNING *;

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
RETURNING id;

-- name: UpdateFlight :one
UPDATE flights SET
number = COALESCE(sqlc.narg('number'), number),
origin_airport_id = COALESCE(sqlc.narg('origin_airport_id'), origin_airport_id),
destination_airport_id = COALESCE(sqlc.narg('destination_airport_id'), destination_airport_id),
published = COALESCE(sqlc.narg('published'), published)
WHERE id=sqlc.arg('id')
RETURNING id;

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

------------------------------------------------------------------------------- routes

-- name: GetRouteByIATACodes :one
SELECT * FROM routes
WHERE origin_airport_iata_code=:origin_airport_iata_code AND destination_airport_iata_code=:destination_airport_iata_code
LIMIT 1;

-- name: ListRoutes :many
SELECT * FROM routes
ORDER BY flights_count DESC, origin_airport_id ASC, destination_airport_id ASC;

-- name: ListFlightsByRoute :many
SELECT *
FROM flights_view
WHERE origin_airport_id=:origin_airport OR destination_airport_id=:destination_airport
ORDER BY id ASC;