CREATE TABLE if NOT EXISTS bookcase (
    id serial PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE if NOT EXISTS books (
    id serial PRIMARY KEY,
    author text NOT NULL,
    publisher text NOT NULL,
    name text NOT NULL,
    bookcase_id  integer NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS bookcase;
DROP TABLE IF EXISTS books;
