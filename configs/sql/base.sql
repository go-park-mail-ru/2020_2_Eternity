create table users(
	id serial primary key,
	username text unique not null,
	email text unique not null,
	password text not null,
	birthdate date,
	reg_date timestamp,
	avatar text
);

create table pins(
	pin_id serial primary key,
	user_id integer not null,
	data text not null,
	foreign key(user_id) references users(id)
);