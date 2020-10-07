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
	id serial primary key,
	title varchar(255) not null,
	content text not null,
	img_link varchar(255) not null,
	user_id integer not null,
	foreign key(user_id) references users(id)
);
