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
	user_id integer not null,
	data text,
	foreign key(user_id) references users(id)
);