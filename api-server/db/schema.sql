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

CREATE TABLE IF NOT EXISTS airports (
  id INTEGER PRIMARY KEY,
  iata_code TEXT NOT NULL UNIQUE,
  oadb_id INTEGER
);

CREATE TABLE IF NOT EXISTS flight_schedules (
  id INTEGER PRIMARY KEY,
  airline_id INTEGER NOT NULL,
  number TEXT NOT NULL,
  origin_airport_id INTEGER NOT NULL,
  destination_airport_id INTEGER NOT NULL,
  aircraft_type TEXT NOT NULL,
  start_localdate TEXT NOT NULL,
  end_localdate TEXT NOT NULL,
  days_of_week TEXT NOT NULL,
  departure_localtime TEXT NOT NULL,
  duration_sec INTEGER NOT NULL,
  published BOOLEAN NOT NULL,
  FOREIGN KEY (airline_id) REFERENCES airlines(id),
  FOREIGN KEY (origin_airport_id) REFERENCES airports(id),
  FOREIGN KEY (destination_airport_id) REFERENCES airports(id)
);

CREATE VIEW IF NOT EXISTS flight_schedules_view AS
SELECT flight_schedules.*,
  airlines.iata_code AS airline_iata_code,
  airlines.name AS airline_name,
  origin_airport.iata_code AS origin_airport_iata_code,
  origin_airport.oadb_id AS origin_airport_oadb_id,
  destination_airport.iata_code AS destination_airport_iata_code,
  destination_airport.oadb_id AS destination_airport_oadb_id
FROM flight_schedules
JOIN airlines airlines ON airlines.id=flight_schedules.airline_id
JOIN airports origin_airport ON origin_airport.id=flight_schedules.origin_airport_id
JOIN airports destination_airport ON destination_airport.id=flight_schedules.destination_airport_id;

CREATE TABLE IF NOT EXISTS flight_instances (
  id INTEGER PRIMARY KEY,
  source_flight_schedule_id INTEGER,
  source_flight_schedule_instance_localdate TEXT,
  airline_id INTEGER NOT NULL,
  number TEXT NOT NULL,
  origin_airport_id INTEGER NOT NULL,
  destination_airport_id INTEGER NOT NULL,
  aircraft_type TEXT NOT NULL,
  aircraft_id INTEGER,
  departure_datetime TEXT NOT NULL,
  arrival_datetime TEXT NOT NULL,
  departure_datetime_utc DATETIME NOT NULL,
  arrival_datetime_utc DATETIME NOT NULL,
  notes TEXT NOT NULL,
  published BOOLEAN NOT NULL,
  FOREIGN KEY (source_flight_schedule_id) REFERENCES flight_schedules(id),
  FOREIGN KEY (airline_id) REFERENCES airlines(id),
  FOREIGN KEY (origin_airport_id) REFERENCES airports(id),
  FOREIGN KEY (destination_airport_id) REFERENCES airports(id),
  FOREIGN KEY (aircraft_id) REFERENCES aircraft(id)
);

CREATE VIEW IF NOT EXISTS flight_instances_view AS
SELECT flight_instances.*,
  airlines.iata_code AS airline_iata_code,
  airlines.name AS airline_name,
  origin_airport.iata_code AS origin_airport_iata_code,
  origin_airport.oadb_id AS origin_airport_oadb_id,
  destination_airport.iata_code AS destination_airport_iata_code,
  destination_airport.oadb_id AS destination_airport_oadb_id,
  aircraft.registration AS aircraft_registration,
  aircraft.aircraft_type AS aircraft_aircraft_type,
  aircraft.airline_id AS aircraft_airline_id,
  aircraft.airline_iata_code AS aircraft_airline_iata_code,
  aircraft.airline_name AS aircraft_airline_name
FROM flight_instances
JOIN airlines ON airlines.id=flight_instances.airline_id
JOIN airports origin_airport ON origin_airport.id=flight_instances.origin_airport_id
JOIN airports destination_airport ON destination_airport.id=flight_instances.destination_airport_id
LEFT JOIN aircraft_view aircraft ON aircraft.id=flight_instances.aircraft_id;

CREATE UNIQUE INDEX IF NOT EXISTS unique_flight_instance_for_date_in_flight_schedule
ON flight_instances(source_flight_schedule_id, source_flight_schedule_instance_localdate)
WHERE source_flight_schedule_id IS NOT NULL AND source_flight_schedule_instance_localdate IS NOT NULL;

CREATE VIEW IF NOT EXISTS routes AS
SELECT origin_airport_id, destination_airport_id,
  origin_airport_iata_code, origin_airport_oadb_id,
  destination_airport_iata_code, destination_airport_oadb_id,
  COUNT(*) AS flight_schedules_count
FROM flight_schedules_view
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
  flight_instance_id INTEGER NOT NULL,
  PRIMARY KEY (itinerary_id, flight_instance_id),
  FOREIGN KEY (itinerary_id) REFERENCES itineraries(id),
  FOREIGN KEY (flight_instance_id) REFERENCES flight_instances(id)
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
  flight_instance_id INTEGER NOT NULL,
  seat TEXT NOT NULL,
  UNIQUE (itinerary_id, passenger_id, flight_instance_id),
  UNIQUE (flight_instance_id, seat),
  FOREIGN KEY (itinerary_id) REFERENCES itineraries(id),
  FOREIGN KEY (passenger_id) REFERENCES passengers(id),
  FOREIGN KEY (flight_instance_id) REFERENCES flight_instances(id)
);

CREATE VIEW IF NOT EXISTS seat_assignments_view AS
SELECT seat_assignments.*,
  itineraries.record_id AS itinerary_record_id,
  passengers.name AS passenger_name
FROM seat_assignments
JOIN itineraries ON itineraries.id=seat_assignments.itinerary_id
JOIN passengers ON passengers.id=seat_assignments.passenger_id;
