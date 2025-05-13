CREATE TABLE IF NOT EXISTS aircraft (
  id INTEGER PRIMARY KEY,
  registration TEXT NOT NULL UNIQUE,
  aircraft_type TEXT NOT NULL,
  airline_id INTEGER NOT NULL,
  FOREIGN KEY (airline_id) REFERENCES airlines(id)
);

CREATE TABLE IF NOT EXISTS airlines (
  id INTEGER PRIMARY KEY,
  iata_code TEXT NOT NULL UNIQUE,
  name TEXT NOT NULL
);

CREATE VIEW IF NOT EXISTS aircraft_view AS
SELECT aircraft.*, airlines.iata_code AS airline_iata_code, airlines.name AS airline_name
FROM aircraft
JOIN airlines ON airlines.id=aircraft.airline_id;

CREATE TABLE IF NOT EXISTS fleets (
  id INTEGER PRIMARY KEY,
  airline_id INTEGER NOT NULL,
  code TEXT NOT NULL,
  description TEXT NOT NULL,
  UNIQUE (code, airline_id),
  FOREIGN KEY (airline_id) REFERENCES airlines(id)
);

CREATE VIEW IF NOT EXISTS fleets_view AS
SELECT fleets.*, airlines.iata_code AS airline_iata_code, airlines.name AS airline_name
FROM fleets
JOIN airlines ON airlines.id=fleets.airline_id;

CREATE TABLE IF NOT EXISTS fleets_aircraft (
  fleet_id INTEGER NOT NULL,
  aircraft_id INTEGER NOT NULL,
  PRIMARY KEY (fleet_id, aircraft_id),
  FOREIGN KEY (fleet_id) REFERENCES fleets(id),
  FOREIGN KEY (aircraft_id) REFERENCES aircraft(id)
);

CREATE TABLE IF NOT EXISTS airports (
  id INTEGER PRIMARY KEY,
  iata_code TEXT NOT NULL UNIQUE,
  oadb_id INTEGER
);

CREATE TABLE IF NOT EXISTS schedules (
  id INTEGER PRIMARY KEY,
  airline_id INTEGER NOT NULL,
  number TEXT NOT NULL,
  origin_airport_id INTEGER NOT NULL,
  destination_airport_id INTEGER NOT NULL,
  fleet_id INTEGER NOT NULL,
  start_localdate TEXT NOT NULL,
  end_localdate TEXT NOT NULL,
  days_of_week TEXT NOT NULL,
  departure_localtime TEXT NOT NULL,
  duration_sec INTEGER NOT NULL,
  published BOOLEAN NOT NULL,
  FOREIGN KEY (airline_id) REFERENCES airlines(id),
  FOREIGN KEY (fleet_id) REFERENCES fleets(id),
  FOREIGN KEY (origin_airport_id) REFERENCES airports(id),
  FOREIGN KEY (destination_airport_id) REFERENCES airports(id)
);

CREATE VIEW IF NOT EXISTS schedules_view AS
SELECT schedules.*,
  airlines.iata_code AS airline_iata_code,
  airlines.name AS airline_name,
  fleets.airline_id AS fleet_airline_id,
  fleets.code AS fleet_code,
  fleets.description AS fleet_description,
  origin_airport.iata_code AS origin_airport_iata_code,
  origin_airport.oadb_id AS origin_airport_oadb_id,
  destination_airport.iata_code AS destination_airport_iata_code,
  destination_airport.oadb_id AS destination_airport_oadb_id
FROM schedules
JOIN airlines airlines ON airlines.id=schedules.airline_id
JOIN fleets ON fleets.id=schedules.fleet_id
JOIN airports origin_airport ON origin_airport.id=schedules.origin_airport_id
JOIN airports destination_airport ON destination_airport.id=schedules.destination_airport_id;

CREATE TABLE IF NOT EXISTS flights (
  id INTEGER PRIMARY KEY,
  source_schedule_id INTEGER,
  source_schedule_instance_localdate TEXT,
  airline_id INTEGER NOT NULL,
  number TEXT NOT NULL,
  origin_airport_id INTEGER NOT NULL,
  destination_airport_id INTEGER NOT NULL,
  fleet_id INTEGER NOT NULL,
  aircraft_id INTEGER,
  departure_datetime TEXT NOT NULL,
  arrival_datetime TEXT NOT NULL,
  departure_datetime_utc DATETIME NOT NULL,
  arrival_datetime_utc DATETIME NOT NULL,
  notes TEXT NOT NULL,
  published BOOLEAN NOT NULL,
  mileage_reward INTEGER NOT NULL DEFAULT 0,
  FOREIGN KEY (source_schedule_id) REFERENCES schedules(id),
  FOREIGN KEY (airline_id) REFERENCES airlines(id),
  FOREIGN KEY (fleet_id) REFERENCES fleets(id),
  FOREIGN KEY (origin_airport_id) REFERENCES airports(id),
  FOREIGN KEY (destination_airport_id) REFERENCES airports(id),
  FOREIGN KEY (aircraft_id) REFERENCES aircraft(id)
);

CREATE VIEW IF NOT EXISTS flights_view AS
SELECT flights.*,
  airlines.iata_code AS airline_iata_code,
  airlines.name AS airline_name,
  fleets.airline_id AS fleet_airline_id,
  fleets.code AS fleet_code,
  fleets.description AS fleet_description,
  origin_airport.iata_code AS origin_airport_iata_code,
  origin_airport.oadb_id AS origin_airport_oadb_id,
  destination_airport.iata_code AS destination_airport_iata_code,
  destination_airport.oadb_id AS destination_airport_oadb_id,
  aircraft.registration AS aircraft_registration,
  aircraft.aircraft_type AS aircraft_aircraft_type,
  aircraft.airline_id AS aircraft_airline_id,
  aircraft.airline_iata_code AS aircraft_airline_iata_code,
  aircraft.airline_name AS aircraft_airline_name
FROM flights
JOIN airlines ON airlines.id=flights.airline_id
JOIN fleets ON fleets.id=flights.fleet_id
JOIN airports origin_airport ON origin_airport.id=flights.origin_airport_id
JOIN airports destination_airport ON destination_airport.id=flights.destination_airport_id
LEFT JOIN aircraft_view aircraft ON aircraft.id=flights.aircraft_id;

CREATE UNIQUE INDEX IF NOT EXISTS unique_flight_for_date_in_schedule
ON flights(source_schedule_id, source_schedule_instance_localdate)
WHERE source_schedule_id IS NOT NULL AND source_schedule_instance_localdate IS NOT NULL;

CREATE VIEW IF NOT EXISTS routes AS
SELECT origin_airport_id, destination_airport_id,
  origin_airport_iata_code, origin_airport_oadb_id,
  destination_airport_iata_code, destination_airport_oadb_id,
  COUNT(*) AS schedules_count
FROM schedules_view
GROUP BY origin_airport_id, destination_airport_id;

CREATE TABLE IF NOT EXISTS passengers (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS itineraries (
  id INTEGER PRIMARY KEY,
  record_id TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS itinerary_flights (
  itinerary_id INTEGER NOT NULL,
  flight_id INTEGER NOT NULL,
  PRIMARY KEY (itinerary_id, flight_id),
  FOREIGN KEY (itinerary_id) REFERENCES itineraries(id),
  FOREIGN KEY (flight_id) REFERENCES flights(id)
);

CREATE TABLE IF NOT EXISTS itinerary_passengers (
  itinerary_id INTEGER NOT NULL,
  passenger_id INTEGER NOT NULL,
  PRIMARY KEY (itinerary_id, passenger_id),
  FOREIGN KEY (itinerary_id) REFERENCES itineraries(id),
  FOREIGN KEY (passenger_id) REFERENCES passengers(id)
);

CREATE TABLE IF NOT EXISTS seat_assignments (
  id INTEGER PRIMARY KEY,
  itinerary_id INTEGER NOT NULL,
  passenger_id INTEGER NOT NULL,
  flight_id INTEGER NOT NULL,
  seat TEXT NOT NULL,
  UNIQUE (itinerary_id, passenger_id, flight_id),
  UNIQUE (flight_id, seat),
  FOREIGN KEY (itinerary_id) REFERENCES itineraries(id),
  FOREIGN KEY (passenger_id) REFERENCES passengers(id),
  FOREIGN KEY (flight_id) REFERENCES flights(id)
);

CREATE VIEW IF NOT EXISTS seat_assignments_view AS
SELECT seat_assignments.*,
  itineraries.record_id AS itinerary_record_id,
  passengers.name AS passenger_name
FROM seat_assignments
JOIN itineraries ON itineraries.id=seat_assignments.itinerary_id
JOIN passengers ON passengers.id=seat_assignments.passenger_id;

CREATE TABLE IF NOT EXISTS airline_loyalties (
  id INTEGER PRIMARY KEY,
  airline_id INTEGER NOT NULL,
  passenger_id INTEGER NOT NULL,
  mileage_balance INTEGER NOT NULL DEFAULT 0,
  UNIQUE (airline_id, passenger_id),
  FOREIGN KEY (airline_id) REFERENCES airlines(id),
  FOREIGN KEY (passenger_id) REFERENCES passengers(id)
);
