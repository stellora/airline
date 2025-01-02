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
  published BOOLEAN NOT NULL DEFAULT FALSE,
  FOREIGN KEY (airline_id) REFERENCES airlines(id),
  FOREIGN KEY (origin_airport_id) REFERENCES airports(id),
  FOREIGN KEY (destination_airport_id) REFERENCES airports(id)
);

CREATE VIEW IF NOT EXISTS flight_schedules_view AS
SELECT flight_schedules.*,
  origin_airport.iata_code AS origin_airport_iata_code,
  origin_airport.oadb_id AS origin_airport_oadb_id,
  destination_airport.iata_code AS destination_airport_iata_code,
  destination_airport.oadb_id AS destination_airport_oadb_id,
  airlines.iata_code AS airline_iata_code,
  airlines.name AS airline_name
FROM flight_schedules
JOIN airlines airlines ON airlines.id=flight_schedules.airline_id
JOIN airports origin_airport ON origin_airport.id=flight_schedules.origin_airport_id
JOIN airports destination_airport ON destination_airport.id=flight_schedules.destination_airport_id;

CREATE VIEW IF NOT EXISTS routes AS
SELECT origin_airport_id, destination_airport_id,
  origin_airport_iata_code, origin_airport_oadb_id,
  destination_airport_iata_code, destination_airport_oadb_id,
  COUNT(*) AS flight_schedules_count
FROM flight_schedules_view
GROUP BY origin_airport_id, destination_airport_id;