create table if not exists chats(
	id serial primary key,
	creation_time timestamp,

    last_msg_id integer default 0,  -- trigger
	last_msg_content text default '',  -- trigger
	last_msg_username text default '',  -- trigger
	last_msg_time timestamp not null default NOW()  -- trigger
);



create table if not exists uu_chat(
	user_id integer not null,
	collocutor_id integer not null,
	chat_id integer not null,
	last_read_msg_id int default 0,

	new_messages int default 0,  -- trigger

	primary key (user_id, collocutor_id),
	foreign key(user_id) references users(id) on delete cascade,
	foreign key(collocutor_id) references users(id) on delete cascade,
	foreign key(chat_id) references chats(id) on delete cascade
);


create table if not exists messages(
	id serial primary key,
	content text not null,
	creation_time timestamp not null default NOW(),
	chat_id integer not null,
	user_id integer not null,
	is_read bool not null,  -- should de deleted
	foreign key(user_id) references users(id) on delete cascade,
	foreign key(chat_id) references chats(id) on delete cascade
);

create table if not exists msg_files(
	id serial primary key,
	name varchar(127) not null,
	type varchar(127) not null,
	msg_id integer not null,
	foreign key(msg_id) references messages(id) on delete cascade
);




CREATE OR REPLACE FUNCTION msg_change() RETURNS TRIGGER AS $msg_change$
    BEGIN
        IF (TG_OP = 'DELETE') THEN
            IF (OLD.id > (
                    SELECT last_read_msg_id FROM uu_chat
                    where uu_chat.chat_id = OLD.chat_id AND
                    uu_chat.collocutor_id = OLD.user_id)) THEN
                UPDATE uu_chat
                SET new_messages = new_messages - 1
                WHERE uu_chat.chat_id = OLD.chat_id AND
                   uu_chat.collocutor_id = OLD.user_id;
             END IF;

            IF (OLD.id = (SELECT last_msg_id FROM chats where chats.id = OLD.chat_id)) THEN

                CREATE TEMPORARY TABLE last_msg
                ON COMMIT DROP
                AS
                SELECT * FROM messages ORDER BY id DESC LIMIT 1;

                IF ((SELECT COUNT(1) FROM last_msg) = 0) THEN
                    INSERT INTO last_msg (id, content, user_id, creation_time)
                    VALUES (0, '', 0, NOW());
                END IF;

                UPDATE chats
                SET last_msg_id = last_msg.id,
                 last_msg_content = last_msg.content,
	             last_msg_username = (SELECT username FROM users WHERE id = last_msg.user_id),
	             last_msg_time = last_msg.creation_time
	            FROM last_msg
                WHERE chats.id = OLD.chat_id;

            END IF;


        ELSIF (TG_OP = 'UPDATE') THEN
            IF (OLD.id = (SELECT last_msg_id FROM chats where chats.id = OLD.chat_id)) THEN
                UPDATE chats
                SET last_msg_id = NEW.id,
                    last_msg_content = NEW.content,
	                last_msg_username = (SELECT username FROM users WHERE id = NEW.user_id),
	                last_msg_time = NEW.creation_time
                WHERE chats.id = NEW.chat_id;
             END IF;

        ELSIF (TG_OP = 'INSERT') THEN
             UPDATE uu_chat
             SET new_messages = new_messages + 1
             where uu_chat.chat_id = NEW.chat_id AND uu_chat.collocutor_id = NEW.user_id;

             UPDATE uu_chat
             SET last_read_msg_id = NEW.id
             where uu_chat.chat_id = NEW.chat_id AND uu_chat.user_id = NEW.user_id;

             UPDATE chats
             SET last_msg_id = NEW.id,
                 last_msg_content = NEW.content,
	             last_msg_username = (SELECT username FROM users WHERE id = NEW.user_id),
	             last_msg_time = NEW.creation_time
             WHERE chats.id = NEW.chat_id;

        END IF;
        RETURN NULL; -- возвращаемое значение для триггера AFTER игнорируется
    END;
$msg_change$ LANGUAGE plpgsql;


CREATE TRIGGER upd_msgs
AFTER INSERT OR DELETE OR UPDATE ON messages
    FOR EACH ROW EXECUTE PROCEDURE msg_change();


