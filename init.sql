CREATE DATABASE better_world_db;

\connect better_world_db;

CREATE TYPE pin_type AS ENUM ('comparison', 'image', 'video');

CREATE TABLE Pin(
  id SERIAL NOT NULL PRIMARY KEY,
  lat FLOAT( 10, 6 ) NOT NULL,
  lng FLOAT( 10, 6 ) NOT NULL,
  address TEXT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
);

CREATE TABLE PinPhoto(
  id SERIAL NOT NULL PRIMARY KEY,
  type pin_type,
  baseImage TEXT,
  afterImage TEXT,
  FOREIGN KEY(fk_pin_id) REFERENCES Pins(id),
);

CREATE TABLE PinVideo(
  id SERIAL NOT NULL PRIMARY KEY,
  type pin_type,
  video TEXT,
  FOREIGN KEY(fk_pin_id) REFERENCES Pins(id),
)
