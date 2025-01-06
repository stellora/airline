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
WHERE id=sqlc.arg('id')
RETURNING *;

-- name: DeleteAirline :exec
DELETE FROM airlines
WHERE id=?;

-- name: DeleteAllAirlines :exec
DELETE FROM airlines;

------------------------------------------------------------------------------- fleets

-- name: GetFleet :one
SELECT * FROM fleets_view
WHERE id=? LIMIT 1;

-- name: GetFleetByCode :one
SELECT * FROM fleets_view
WHERE airline_id=? AND code=? LIMIT 1;

-- name: ListFleets :many
SELECT * FROM fleets_view
ORDER BY id ASC;

-- name: CreateFleet :one
INSERT INTO fleets (
  airline_id,
  code,
  description
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateFleet :one
UPDATE fleets SET
code = COALESCE(sqlc.narg('code'), code),
description = COALESCE(sqlc.narg('description'), description)
WHERE id=sqlc.arg('id')
RETURNING *;

-- name: DeleteFleet :exec
DELETE FROM fleets
WHERE id=?;

-- name: ListFleetsByAirline :many
SELECT *
FROM fleets_view
WHERE airline_id=:airline
ORDER BY id ASC;

-- name: ListAircraftByFleet :many
SELECT *
FROM aircraft_view
WHERE id IN (
    SELECT aircraft_id
    FROM fleets_aircraft
    WHERE fleet_id = ?
)
ORDER BY id ASC;

-- name: AddAircraftToFleet :exec
INSERT INTO fleets_aircraft (
  fleet_id,
  aircraft_id
) VALUES (
  ?, ?
);

-- name: RemoveAircraftFromFleet :exec
DELETE FROM fleets_aircraft
WHERE fleet_id=? AND aircraft_id=?;

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

-- name: ListSchedulesByAirport :many
SELECT *
FROM schedules_view
WHERE origin_airport_id=:airport OR destination_airport_id=:airport
ORDER BY id ASC;

------------------------------------------------------------------------------- schedules

-- name: GetSchedule :one
SELECT * FROM schedules_view
WHERE id=? LIMIT 1;

-- name: ListSchedules :many
SELECT * FROM schedules_view
ORDER BY id ASC;

-- name: CreateSchedule :one
INSERT INTO schedules (
  airline_id, number, origin_airport_id, destination_airport_id, fleet_id, start_localdate, end_localdate, days_of_week, departure_localtime, duration_sec, published
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING id;

-- name: UpdateSchedule :one
UPDATE schedules SET
number = COALESCE(sqlc.narg('number'), number),
origin_airport_id = COALESCE(sqlc.narg('origin_airport_id'), origin_airport_id),
destination_airport_id = COALESCE(sqlc.narg('destination_airport_id'), destination_airport_id),
fleet_id = COALESCE(sqlc.narg('fleet_id'), fleet_id),
start_localdate = COALESCE(sqlc.narg('start_localdate'), start_localdate),
end_localdate = COALESCE(sqlc.narg('end_localdate'), end_localdate),
days_of_week = COALESCE(sqlc.narg('days_of_week'), days_of_week),
departure_localtime = COALESCE(sqlc.narg('departure_localtime'), departure_localtime),
duration_sec = COALESCE(sqlc.narg('duration_sec'), duration_sec),
published = COALESCE(sqlc.narg('published'), published)
WHERE id=sqlc.arg('id')
RETURNING id;

-- name: DeleteSchedule :exec
DELETE FROM schedules
WHERE id=?;

-- name: DeleteAllSchedules :exec
DELETE FROM schedules;

-- name: ListSchedulesByAirline :many
SELECT *
FROM schedules_view
WHERE airline_id=:airline
ORDER BY id ASC;

-- name: ListSchedulesByRoute :many
SELECT *
FROM schedules_view
WHERE origin_airport_id=:origin_airport AND destination_airport_id=:destination_airport
ORDER BY id ASC;

------------------------------------------------------------------------------- flights

-- name: GetFlight :one
SELECT *
FROM flights_view
WHERE id=sqlc.arg('id') LIMIT 1;

-- name: ListFlights :many
SELECT *
FROM flights_view
ORDER BY departure_datetime_utc ASC, arrival_datetime_utc ASC, id ASC;

-- name: CreateFlight :one
INSERT INTO flights (
  source_schedule_id,
  source_schedule_instance_localdate,
  airline_id,
  number,
  origin_airport_id,
  destination_airport_id,
  fleet_id,
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

-- name: UpdateFlight :one
UPDATE flights SET
number = COALESCE(sqlc.narg('number'), number),
origin_airport_id = COALESCE(sqlc.narg('origin_airport_id'), origin_airport_id),
destination_airport_id = COALESCE(sqlc.narg('destination_airport_id'), destination_airport_id),
fleet_id = COALESCE(sqlc.narg('fleet_id'), fleet_id),
aircraft_id = COALESCE(sqlc.narg('aircraft_id'), aircraft_id),
departure_datetime = COALESCE(sqlc.narg('departure_datetime'), departure_datetime),
arrival_datetime = COALESCE(sqlc.narg('arrival_datetime'), arrival_datetime),
departure_datetime_utc = COALESCE(sqlc.narg('departure_datetime_utc'), departure_datetime_utc),
arrival_datetime_utc = COALESCE(sqlc.narg('arrival_datetime_utc'), arrival_datetime_utc),
notes = COALESCE(sqlc.narg('notes'), notes),
published = COALESCE(sqlc.narg('published'), published)
WHERE id=sqlc.arg('id')
RETURNING id;

-- name: DeleteFlight :exec
DELETE FROM flights
WHERE id=?;

-- name: ListFlightsForSchedule :many
SELECT *
FROM flights_view
WHERE source_schedule_id IS NOT NULL AND source_schedule_id=sqlc.arg('schedule_id')
ORDER BY departure_datetime_utc ASC, arrival_datetime_utc ASC, id ASC;

-- name: ListFlightsByAirline :many
SELECT flights_view.*
FROM flights_view
WHERE airline_id=sqlc.arg('airline_id')
ORDER BY departure_datetime_utc ASC, arrival_datetime_utc ASC, id ASC;

-- name: ListFlightsByRoute :many
SELECT *
FROM flights_view
WHERE origin_airport_id=:origin_airport AND destination_airport_id=:destination_airport
ORDER BY id ASC;

------------------------------------------------------------------------------- passengers

-- name: ListPassengers :many
SELECT * FROM passengers
ORDER BY id ASC;

-- name: GetPassenger :one
SELECT * FROM passengers
WHERE id = ?
LIMIT 1;

-- name: CreatePassenger :one
INSERT INTO passengers (
  name
) VALUES (
  ?
)
RETURNING id;

-- name: UpdatePassenger :one
UPDATE passengers SET
name = COALESCE(sqlc.narg('name'), name)
WHERE id = sqlc.arg('id')
RETURNING id;

-- name: DeletePassenger :exec
DELETE FROM passengers
WHERE id = ?;

------------------------------------------------------------------------------- seat_assignments

-- name: ListSeatAssignmentsForFlight :many
SELECT * FROM seat_assignments_view
WHERE flight_id = ?
ORDER BY id ASC;

-- name: GetSeatAssignment :one
SELECT * FROM seat_assignments_view
WHERE id = ?
LIMIT 1;

-- name: CreateSeatAssignment :one
INSERT INTO seat_assignments (
  itinerary_id,
  passenger_id,
  flight_id,
  seat
) VALUES (
  ?, ?, ?, ?
)
RETURNING id;

-- name: UpdateSeatAssignment :one
UPDATE seat_assignments SET
seat = COALESCE(sqlc.narg('seat'), seat)
WHERE id=sqlc.arg('id')
RETURNING *;

-- name: DeleteSeatAssignment :exec
DELETE FROM seat_assignments
WHERE id=?;

------------------------------------------------------------------------------- itineraries

-- name: ListItineraries :many
SELECT * FROM itineraries
ORDER BY id ASC;

-- name: GetItinerary :one
SELECT * FROM itineraries
WHERE id = ?
LIMIT 1;

-- name: GetItineraryByRecordLocator :one
SELECT * FROM itineraries
WHERE record_id = ?
LIMIT 1;

-- name: CreateItinerary :one
INSERT INTO itineraries (
  record_id
) VALUES (
  ?
)
RETURNING *;

-- name: AddFlightToItinerary :exec
INSERT INTO itinerary_flights (
  itinerary_id,
  flight_id
) VALUES (
  ?, ?
);

-- name: RemoveFlightFromItinerary :exec
DELETE FROM itinerary_flights
WHERE itinerary_id = ? AND flight_id = ?;

-- name: ListItineraryFlights :many
SELECT flights_view.*
FROM flights_view
JOIN itinerary_flights ON itinerary_flights.flight_id = flights_view.id
WHERE itinerary_flights.itinerary_id = sqlc.arg('itinerary_id')
ORDER BY departure_datetime_utc ASC;

-- name: ListItineraryPassengers :many
SELECT passengers.*
FROM passengers
JOIN itinerary_passengers ON itinerary_passengers.passenger_id = passengers.id
WHERE itinerary_passengers.itinerary_id = sqlc.arg('itinerary_id')
ORDER BY passengers.id ASC;

-- name: AddPassengerToItinerary :exec
INSERT INTO itinerary_passengers (
  itinerary_id,
  passenger_id
) VALUES (
  ?, ?
);

-- name: RemovePassengerFromItinerary :exec
DELETE FROM itinerary_passengers
WHERE itinerary_id = ? AND passenger_id = ?;

-- name: DeleteItinerary :exec
DELETE FROM itineraries
WHERE id = ?;


------------------------------------------------------------------------------- routes

-- name: GetRouteByIATACodes :one
SELECT * FROM routes
WHERE origin_airport_iata_code=:origin_airport_iata_code AND destination_airport_iata_code=:destination_airport_iata_code
LIMIT 1;

-- name: ListRoutes :many
SELECT * FROM routes
ORDER BY schedules_count DESC, origin_airport_id ASC, destination_airport_id ASC;