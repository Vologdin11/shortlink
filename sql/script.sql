CREATE DATABASE shortlink;

\c shortlink;

CREATE TABLE shortlink (
    link VARCHAR(500),
    shortlink VARCHAR(50)
);

INSERT INTO shortlink(link, shortlink)
VALUES ('https://google.com', 'http://localhost:8000/gl');