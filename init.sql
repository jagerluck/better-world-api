CREATE DATABASE better_world_db;

\connect better_world_db;

CREATE TYPE pin_types AS ENUM ('comparison', 'image', 'video');

CREATE TABLE Pins(
  id SERIAL NOT NULL PRIMARY KEY,
  lat FLOAT( 10, 6 ) NOT NULL,
  lng FLOAT( 10, 6 ) NOT NULL,
  FOREIGN KEY(pin_media) REFERENCES PinMedia(id),
);

CREATE TABLE PinMedia(
  id SERIAL NOT NULL PRIMARY KEY,
  baseImage TEXT,
  afterImage TEXT,
  video TEXT,
);
