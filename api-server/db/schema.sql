CREATE TABLE IF NOT EXISTS flights (
  id INTEGER PRIMARY KEY,
  number TEXT NOT NULL,
  origin_airport_id INTEGER NOT NULL,
  destination_airport_id INTEGER NOT NULL,
  published BOOLEAN NOT NULL DEFAULT FALSE,
  FOREIGN KEY (origin_airport_id) REFERENCES airports(id),
  FOREIGN KEY (destination_airport_id) REFERENCES airports(id)
);

CREATE TABLE IF NOT EXISTS airports (
  id INTEGER PRIMARY KEY,
  iata_code TEXT NOT NULL UNIQUE,
  oadb_id INTEGER
);

CREATE VIEW IF NOT EXISTS flights_view AS
SELECT flights.*,
  origin_airport.iata_code AS origin_airport_iata_code,
  origin_airport.oadb_id AS origin_airport_oadb_id,
  destination_airport.iata_code AS destination_airport_iata_code,
  destination_airport.oadb_id AS destination_airport_oadb_id
FROM flights
JOIN airports origin_airport ON origin_airport.id=flights.origin_airport_id
JOIN airports destination_airport ON destination_airport.id=flights.destination_airport_id;

CREATE VIEW IF NOT EXISTS routes AS
SELECT origin_airport_id, destination_airport_id,
  origin_airport_iata_code, origin_airport_oadb_id,
  destination_airport_iata_code, destination_airport_oadb_id,
  COUNT(*) AS flights_count
FROM flights_view
GROUP BY origin_airport_id, destination_airport_id;