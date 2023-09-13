/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

/** This is test table. Remove this table and replace with your own tables. */
CREATE TABLE test (
	id serial PRIMARY KEY,
	name VARCHAR ( 50 ) UNIQUE NOT NULL
);

INSERT INTO test (name) VALUES ('test1');
INSERT INTO test (name) VALUES ('test2');


create table users
(
    user_id serial PRIMARY KEY,
    user_phone_number varchar(16) unique,
    user_full_name    varchar(60),
    user_password     varchar(200)                                      not null,
    user_logged       integer default 0                                 not null
);

alter table users
    owner to postgres;

INSERT INTO public.users (user_phone_number, user_full_name, user_password, user_logged) VALUES ('+6285711111112', 'Sawit Pro', '$2a$10$cy5ivLhl8WmHTd9Iq7wVX./.cSP7d/rqQ40Jq3JQ84cEMl8MlKTlm', 5);