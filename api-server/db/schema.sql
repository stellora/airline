CREATE TABLE IF NOT EXISTS flights (
  id INTEGER PRIMARY KEY,
  number TEXT NOT NULL,
  origin_airport INTEGER NOT NULL,
  destination_airport INTEGER NOT NULL,
  published BOOLEAN NOT NULL DEFAULT FALSE,
  FOREIGN KEY (origin_airport) REFERENCES airports(id),
  FOREIGN KEY (destination_airport) REFERENCES airports(id)
);

CREATE TABLE IF NOT EXISTS airports (
  id INTEGER PRIMARY KEY,
  iata_code TEXT NOT NULL,
  oadb_id INTEGER
);
