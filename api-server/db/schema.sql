CREATE TABLE IF NOT EXISTS flights (
  id INTEGER PRIMARY KEY,
  title TEXT NOT NULL,
  published BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS airports (
  id INTEGER PRIMARY KEY,
  title TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS flight_airports (
  flight_id INTEGER NOT NULL,
  airport_id INTEGER NOT NULL,
  PRIMARY KEY (flight_id, airport_id),
  FOREIGN KEY (flight_id) REFERENCES flights(id),
  FOREIGN KEY (airport_id) REFERENCES airports(id)
);