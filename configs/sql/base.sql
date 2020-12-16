create table if not exists users(
	id serial primary key,
	username text unique not null,
	email text unique not null,
	password text not null,
	name text,
	surname text,
	description text,
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

alter table pins add column if not exists height integer default 0;
alter table pins add column if not exists width integer default 0;

create table if not exists follows(
	id1 int not null,
	id2 int not null,
	foreign key (id1) references users(id),
	foreign key (id2) references users(id),
	unique (id1, id2)
);

create table if not exists stats(
	id int unique not null,
	followers int default 0,
	following int default 0,
	foreign key(id) references users(id)
);

CREATE OR REPLACE FUNCTION upd_stats() RETURNS TRIGGER AS $upd_stats$
    BEGIN
        IF (TG_OP = 'DELETE') THEN
            UPDATE stats set following = following - 1 where stats.id = OLD.id1;
			UPDATE stats set followers = followers - 1 where stats.id = OLD.id2;
            RETURN OLD;
        ELSIF (TG_OP = 'INSERT') THEN
            UPDATE stats set following = following + 1 where stats.id = NEW.id1;
			UPDATE stats set followers = followers + 1 where stats.id = NEW.id2;
            RETURN NEW;
        END IF;
        RETURN NULL; -- возвращаемое значение для триггера AFTER игнорируется
    END;
$upd_stats$ LANGUAGE plpgsql;

CREATE TRIGGER upd_stats_trg
AFTER INSERT OR DELETE ON follows
    FOR EACH ROW EXECUTE PROCEDURE upd_stats();

CREATE OR REPLACE FUNCTION ins_stats() RETURNS TRIGGER AS $ins_stats$
    BEGIN
        IF (TG_OP = 'DELETE') THEN
            delete from stats where old.id = id;
            RETURN OLD;
        ELSIF (TG_OP = 'INSERT') THEN
			insert into stats(id) values(new.id);
            RETURN NEW;
        END IF;
        RETURN NULL; -- возвращаемое значение для триггера AFTER игнорируется
    END;
$ins_stats$ LANGUAGE plpgsql;

CREATE TRIGGER ins_stats_trg
AFTER INSERT OR DELETE ON users
    FOR EACH ROW EXECUTE PROCEDURE ins_stats();

create table if not exists boards(
	id serial primary key,
	title varchar(255) not null,
	content text not null,
	user_id integer not null,
	foreign key(user_id) references users(id)
);

create table if not exists boards_pins(
	board_id int not null,
	pin_id int not null,
	foreign key (board_id) references boards(id),
	foreign key (pin_id) references pins(id),
	unique (board_id, pin_id)
);


create table pins_vectors_content (
    idv int unique not null,
    vec tsvector,
	foreign key(idv) references pins(id)
);

CREATE INDEX idx_gin_pins_content
ON pins_vectors_content
USING gin ("vec");

CREATE OR REPLACE FUNCTION ins_pin_vct_con() RETURNS TRIGGER AS $ins_pin_vct_con$
    BEGIN
        IF (TG_OP = 'UPDATE') THEN
            update pins_vectors_content set vec=to_tsvector(new.content) where old.idv = idv;
            RETURN OLD;
        ELSIF (TG_OP = 'INSERT') THEN
			insert into pins_vectors_content(idv, vec) values(new.id, to_tsvector(new.content));
            RETURN NEW;
        END IF;
        RETURN NULL; -- возвращаемое значение для триггера AFTER игнорируется
    END;
$ins_pin_vct_con$ LANGUAGE plpgsql;

CREATE TRIGGER ins_pin_vct_con_trg
AFTER INSERT OR UPDATE ON pins
    FOR EACH ROW EXECUTE PROCEDURE ins_pin_vct_con();

insert into pins_vectors_content(idv, vec) select id, to_tsvector(content) from pins on conflict do nothing;

create table boards_vectors_content (
    idv int unique not null,
    vec tsvector,
	foreign key(idv) references boards(id)
);

CREATE INDEX idx_gin_boards_content
ON boards_vectors_content
USING gin ("vec");

CREATE OR REPLACE FUNCTION ins_board_vct_con() RETURNS TRIGGER AS $ins_board_vct_con$
    BEGIN
        IF (TG_OP = 'UPDATE') THEN
            update boards_vectors_content set vec=to_tsvector(new.content) where old.idv = idv;
            RETURN OLD;
        ELSIF (TG_OP = 'INSERT') THEN
			insert into boards_vectors_content(idv, vec) values(new.id, to_tsvector(new.content));
            RETURN NEW;
        END IF;
        RETURN NULL; -- возвращаемое значение для триггера AFTER игнорируется
    END;
$ins_board_vct_con$ LANGUAGE plpgsql;

CREATE TRIGGER ins_board_vct_con_trg
AFTER INSERT OR UPDATE ON boards
    FOR EACH ROW EXECUTE PROCEDURE ins_board_vct_con();

insert into boards_vectors_content(idv, vec) select id, to_tsvector(content) from boards on conflict do nothing;