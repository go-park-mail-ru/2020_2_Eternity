CREATE TABLE if NOT EXISTS comments (
    id serial primary key,
    path integer[] not null,
    content varchar(200) not null,
    pin_id integer not null,
    user_id integer not null,
    foreign key(pin_id) references pins(id) on delete cascade,
    foreign key(user_id) references users(id) on delete cascade
);

-- CREATE FUNCTION on_delete_comments() RETURNS trigger AS $on_delete_comments$
--     BEGIN
--         DELETE FROM comments
--         WHERE path && ARRAY(SELECT id FROM old_table);
--
--         RETURN NULL;
--     END;
-- $on_delete_comments$ LANGUAGE plpgsql;
--
--
-- CREATE TRIGGER on_delete_from_comments
-- AFTER DELETE
-- ON comments
-- REFERENCING OLD TABLE AS old_table
-- FOR EACH statement
-- EXECUTE PROCEDURE on_delete_comments();