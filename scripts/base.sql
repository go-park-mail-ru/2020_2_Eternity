create database eternity
	with owner postgres
	encoding 'utf8'
	LC_COLLATE = 'ru_RU.UTF-8'
    LC_CTYPE = 'ru_RU.UTF-8'
    TABLESPACE = pg_default
	;
drop table if exists users cascade;
drop table if exists pins cascade;

create table users(
	id serial primary key,
	username text unique not null,
	email text unique not null,
	password text not null,
	age integer,
	reg_date timestamp
);

create table pins(
	pin_id serial primary key,
	user_id integer,
	data text,
	foreign key(user_id) references users(id)
);