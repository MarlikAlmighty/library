CREATE TABLE if NOT EXISTS shelfs (
	shelfs_id serial PRIMARY KEY,
	shelfs_name text NOT NULL,
	shelfs_number integer NOT NULL
    );

CREATE TABLE if NOT EXISTS books (
	books_id serial PRIMARY KEY,
	author text NOT NULL,
	num_pages integer NOT NULL,
	publisher text NOT NULL,
	shelfs_id  integer NOT NULL,
    pages_id integer NOT NULL
    );

CREATE TABLE if NOT EXISTS pages (
	pages_id serial PRIMARY KEY,
	books_id integer NOT NULL,
	pages_number  integer NOT NULL,
    pages_text text NOT NULL
    );
