CREATE DATABASE shortlink;

\c shortlink;

CREATE TABLE shortlink (
    link VARCHAR(500),
    shortlink VARCHAR(50)
);