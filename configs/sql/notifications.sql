create table if not exists notifications(
	id serial primary key,
	to_user_id integer not null,
	type integer not null,
	encoded_data bytea,
	creation_time timestamp,
	is_read bool,
	foreign key(to_user_id) references users(id) on delete cascade
);
