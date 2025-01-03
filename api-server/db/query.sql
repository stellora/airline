------------------------------------------------------------------------------- aircraft

-- name: GetAircraft :one
SELECT * FROM aircraft_view
WHERE id=? LIMIT 1;

-- name: GetAircraftByRegistration :one
SELECT * FROM aircraft_view
WHERE registration=? LIMIT 1;

-- name: ListAircraft :many
SELECT * FROM aircraft_view
ORDER BY id ASC;

-- name: CreateAircraft :one
INSERT INTO aircraft (
  registration,
  aircraft_type,
  airline_id
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateAircraft :one
UPDATE aircraft SET
registration = COALESCE(sqlc.narg('registration'), registration),
aircraft_type = COALESCE(sqlc.narg('aircraft_type'), aircraft_type),
airline_id = COALESCE(sqlc.narg('airline_id'), airline_id)
WHERE id=sqlc.arg('id')
RETURNING *;

-- name: DeleteAircraft :exec
DELETE FROM aircraft
WHERE id=?;

-- name: DeleteAllAircraft :exec
DELETE FROM aircraft;

-- name: ListAircraftByAirline :many
SELECT *
FROM aircraft_view
WHERE airline_id=:airline
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
iata_code = COALESCE(sqlc.narg('iata_code'), iata_code),
oadb_id = COALESCE(sqlc.narg('oadb_id'), oadb_id)
WHERE id=sqlc.arg('id')
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

------------------------------------------------------------------------------- flight_schedules

-- name: GetFlightSchedule :one
SELECT * FROM flight_schedules_view
WHERE id=? LIMIT 1;

-- name: ListFlightSchedules :many
SELECT * FROM flight_schedules_view
ORDER BY id ASC;

-- name: CreateFlightSchedule :one
INSERT INTO flight_schedules (
  airline_id, number, origin_airport_id, destination_airport_id, aircraft_type, start_date, end_date, days_of_week, published
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING id;

-- name: UpdateFlightSchedule :one
UPDATE flight_schedules SET
airline_id = COALESCE(sqlc.narg('airline_id'), airline_id),
number = COALESCE(sqlc.narg('number'), number),
origin_airport_id = COALESCE(sqlc.narg('origin_airport_id'), origin_airport_id),
destination_airport_id = COALESCE(sqlc.narg('destination_airport_id'), destination_airport_id),
aircraft_type = COALESCE(sqlc.narg('aircraft_type'), aircraft_type),
start_date = COALESCE(sqlc.narg('start_date'), start_date),
end_date = COALESCE(sqlc.narg('end_date'), end_date),
days_of_week = COALESCE(sqlc.narg('days_of_week'), days_of_week),
published = COALESCE(sqlc.narg('published'), published)
WHERE id=sqlc.arg('id')
RETURNING id;

-- name: DeleteFlightSchedule :exec
DELETE FROM flight_schedules
WHERE id=?;

-- name: DeleteAllFlightSchedules :exec
DELETE FROM flight_schedules;

-- name: ListFlightSchedulesByAirline :many
SELECT *
FROM flight_schedules_view
WHERE airline_id=:airline
ORDER BY id ASC;

-- name: ListFlightSchedulesByRoute :many
SELECT *
FROM flight_schedules_view
WHERE origin_airport_id=:origin_airport OR destination_airport_id=:destination_airport
ORDER BY id ASC;

------------------------------------------------------------------------------- flight_instances

-- name: GetFlightInstance :one
SELECT sqlc.embed(flight_instances), sqlc.embed(flight_schedules_view)
FROM flight_instances
JOIN flight_schedules_view ON flight_instances.source_flight_schedule_id=flight_schedules_view.id
WHERE flight_instances.id=sqlc.arg('id') LIMIT 1;

-- name: ListFlightInstances :many
SELECT sqlc.embed(flight_instances), sqlc.embed(flight_schedules_view)
FROM flight_instances
JOIN flight_schedules_view ON flight_instances.source_flight_schedule_id=flight_schedules_view.id
ORDER BY flight_instances.id ASC;

-- name: CreateFlightInstance :one
INSERT INTO flight_instances (
  source_flight_schedule_id,
  instance_date,
  aircraft_id,
  notes
) VALUES (
  ?, ?, ?, ?
)
RETURNING id;

-- name: UpdateFlightInstance :one
UPDATE flight_instances SET
aircraft_id = COALESCE(sqlc.narg('aircraft_id'), aircraft_id),
notes = COALESCE(sqlc.narg('notes'), notes)
WHERE id=sqlc.arg('id')
RETURNING id;

-- name: DeleteFlightInstance :exec
DELETE FROM flight_schedules
WHERE id=?;

-- name: ListFlightInstancesForFlightSchedule :many
SELECT sqlc.embed(flight_instances), sqlc.embed(flight_schedules_view)
FROM flight_instances
JOIN flight_schedules_view ON flight_instances.source_flight_schedule_id=flight_schedules_view.id
WHERE source_flight_schedule_id=sqlc.arg('flight_schedule_id')
ORDER BY flight_instances.id ASC;

------------------------------------------------------------------------------- routes

-- name: GetRouteByIATACodes :one
SELECT * FROM routes
WHERE origin_airport_iata_code=:origin_airport_iata_code AND destination_airport_iata_code=:destination_airport_iata_code
LIMIT 1;

-- name: ListRoutes :many
SELECT * FROM routes
ORDER BY flight_schedules_count DESC, origin_airport_id ASC, destination_airport_id ASC;