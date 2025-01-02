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

-- name: ListFlightSchedulesByAirport :many
SELECT *
FROM flight_schedules_view
WHERE origin_airport_id=:airport OR destination_airport_id=:airport
ORDER BY id ASC;

------------------------------------------------------------------------------- airlines

-- name: GetAirline :one
SELECT * FROM airlines
WHERE id=? LIMIT 1;

-- name: GetAirlineByIATACode :one
SELECT * FROM airlines
WHERE iata_code=? LIMIT 1;

-- name: ListAirlines :many
SELECT * FROM airlines
ORDER BY id ASC;

-- name: CreateAirline :one
INSERT INTO airlines (
  iata_code,
  name
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateAirline :one
UPDATE airlines SET
iata_code = COALESCE(sqlc.narg('iata_code'), iata_code),
name = COALESCE(sqlc.narg('name'), name)
WHERE id=?
RETURNING *;

-- name: DeleteAirline :exec
DELETE FROM airlines
WHERE id=?;

-- name: DeleteAllAirlines :exec
DELETE FROM airlines;

-- name: ListFlightSchedulesByAirline :many
SELECT *
FROM flight_schedules_view
WHERE airline_id=:airline
ORDER BY id ASC;

------------------------------------------------------------------------------- flight_schedules

-- name: GetFlightSchedule :one
SELECT * FROM flight_schedules_view
WHERE id=? LIMIT 1;

-- name: ListFlightSchedules :many
SELECT * FROM flight_schedules_view
ORDER BY id ASC;

-- name: CreateFlightSchedule :one
INSERT INTO flight_schedules (
  airline_id, number, origin_airport_id, destination_airport_id, published
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING id;

-- name: UpdateFlightSchedule :one
UPDATE flight_schedules SET
airline_id = COALESCE(sqlc.narg('airline_id'), airline_id),
number = COALESCE(sqlc.narg('number'), number),
origin_airport_id = COALESCE(sqlc.narg('origin_airport_id'), origin_airport_id),
destination_airport_id = COALESCE(sqlc.narg('destination_airport_id'), destination_airport_id),
published = COALESCE(sqlc.narg('published'), published)
WHERE id=sqlc.arg('id')
RETURNING id;

-- name: DeleteFlightSchedule :exec
DELETE FROM flight_schedules
WHERE id=?;

-- name: DeleteAllFlightSchedules :exec
DELETE FROM flight_schedules;

------------------------------------------------------------------------------- routes

-- name: GetRouteByIATACodes :one
SELECT * FROM routes
WHERE origin_airport_iata_code=:origin_airport_iata_code AND destination_airport_iata_code=:destination_airport_iata_code
LIMIT 1;

-- name: ListRoutes :many
SELECT * FROM routes
ORDER BY flight_schedules_count DESC, origin_airport_id ASC, destination_airport_id ASC;

-- name: ListFlightSchedulesByRoute :many
SELECT *
FROM flight_schedules_view
WHERE origin_airport_id=:origin_airport OR destination_airport_id=:destination_airport
ORDER BY id ASC;