create table if not exists notifications(
	id serial primary key,
	from_user_id integer not null,
	to_user_id integer not null,
	type integer not null,
	encoded_data bytea,
	creation_time timestamp,
	foreign key(from_user_id) references users(id) on delete cascade,
	foreign key(to_user_id) references users(id) on delete cascade
);
