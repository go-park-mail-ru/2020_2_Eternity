create table if not exists users(
	id serial primary key,
	username text unique not null,
	email text unique not null,
	password text not null,
	birthdate date,
	reg_date timestamp,
	avatar text
);

create table if not exists pins(
	id serial primary key,
	title varchar(255) not null,
	content text not null,
	user_id integer not null,
	foreign key(user_id) references users(id)
);

create table if not exists pin_images(
	id serial primary key,
	name varchar(127) not null,
	pin_id integer not null,
	foreign key(pin_id) references pins(id) on delete cascade
);
