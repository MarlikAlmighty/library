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

insert into bookcase (id, name) values (1, 'Fantasy');
insert into bookcase (id, name) values (2, 'Horror');
insert into bookcase (id, name) values (3, 'Detectives');
insert into books (id, author, publisher, name, bookcase_id) values (1, 'Lyman Baum', 'www', 'The Wonderful Wizard of Oz', 1);
insert into books (id, author, publisher, name, bookcase_id) values (2, 'Nora Ephron', 'www', 'I Feel Bad About My Neck', 2);
insert into books (id, author, publisher, name, bookcase_id) values (3, 'Alain Mabanckou', 'www', 'Broken Glass', 3);

---- create above / drop below ----

DROP TABLE IF EXISTS bookcase;
DROP TABLE IF EXISTS books;
