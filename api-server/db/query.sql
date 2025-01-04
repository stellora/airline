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
  airline_id, number, origin_airport_id, destination_airport_id, aircraft_type, start_localdate, end_localdate, days_of_week, departure_localtime, arrival_localtime, published
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING id;

-- name: UpdateFlightSchedule :one
UPDATE flight_schedules SET
airline_id = COALESCE(sqlc.narg('airline_id'), airline_id),
number = COALESCE(sqlc.narg('number'), number),
origin_airport_id = COALESCE(sqlc.narg('origin_airport_id'), origin_airport_id),
destination_airport_id = COALESCE(sqlc.narg('destination_airport_id'), destination_airport_id),
aircraft_type = COALESCE(sqlc.narg('aircraft_type'), aircraft_type),
start_localdate = COALESCE(sqlc.narg('start_localdate'), start_localdate),
end_localdate = COALESCE(sqlc.narg('end_localdate'), end_localdate),
days_of_week = COALESCE(sqlc.narg('days_of_week'), days_of_week),
departure_localtime = COALESCE(sqlc.narg('departure_localtime'), departure_localtime),
arrival_localtime = COALESCE(sqlc.narg('arrival_localtime'), arrival_localtime),
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
SELECT *
FROM flight_instances_view
WHERE id=sqlc.arg('id') LIMIT 1;

-- name: ListFlightInstances :many
SELECT *
FROM flight_instances_view
ORDER BY departure_datetime_utc ASC, arrival_datetime_utc ASC, id ASC;

-- name: CreateFlightInstance :one
INSERT INTO flight_instances (
  source_flight_schedule_id,
  source_flight_schedule_instance_localdate,
  airline_id,
  number,
  origin_airport_id,
  destination_airport_id,
  aircraft_type,
  aircraft_id,
  departure_datetime,
  arrival_datetime,
  departure_datetime_utc,
  arrival_datetime_utc,
  notes,
  published
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING id;

-- name: UpdateFlightInstance :one
UPDATE flight_instances SET
airline_id = COALESCE(sqlc.narg('airline_id'), airline_id),
number = COALESCE(sqlc.narg('number'), number),
origin_airport_id = COALESCE(sqlc.narg('origin_airport_id'), origin_airport_id),
destination_airport_id = COALESCE(sqlc.narg('destination_airport_id'), destination_airport_id),
aircraft_type = COALESCE(sqlc.narg('aircraft_type'), aircraft_type),
aircraft_id = COALESCE(sqlc.narg('aircraft_id'), aircraft_id),
departure_datetime = COALESCE(sqlc.narg('departure_datetime'), departure_datetime),
arrival_datetime = COALESCE(sqlc.narg('arrival_datetime'), arrival_datetime),
departure_datetime_utc = COALESCE(sqlc.narg('departure_datetime_utc'), departure_datetime_utc),
arrival_datetime_utc = COALESCE(sqlc.narg('arrival_datetime_utc'), arrival_datetime_utc),
notes = COALESCE(sqlc.narg('notes'), notes),
published = COALESCE(sqlc.narg('published'), published)
WHERE id=sqlc.arg('id')
RETURNING id;

-- name: DeleteFlightInstance :exec
DELETE FROM flight_instances
WHERE id=?;

-- name: ListFlightInstancesForFlightSchedule :many
SELECT *
FROM flight_instances_view
WHERE source_flight_schedule_id IS NOT NULL AND source_flight_schedule_id=sqlc.arg('flight_schedule_id')
ORDER BY departure_datetime_utc ASC, arrival_datetime_utc ASC, id ASC;

------------------------------------------------------------------------------- routes

-- name: GetRouteByIATACodes :one
SELECT * FROM routes
WHERE origin_airport_iata_code=:origin_airport_iata_code AND destination_airport_iata_code=:destination_airport_iata_code
LIMIT 1;

-- name: ListRoutes :many
SELECT * FROM routes
ORDER BY flight_schedules_count DESC, origin_airport_id ASC, destination_airport_id ASC;