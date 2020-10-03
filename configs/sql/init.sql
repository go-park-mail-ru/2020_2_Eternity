CREATE USER pinterest_user WITH password '662f2710-4e08-4be7-a278-a53ae86ba7f6';
create database pinterest_db
	with owner pinterest_user
	encoding 'utf8'
	LC_COLLATE = 'ru_RU.UTF-8'
    LC_CTYPE = 'ru_RU.UTF-8'
    TABLESPACE = pg_default
	;
GRANT ALL PRIVILEGES ON database pinterest_db TO pinterest_user;
