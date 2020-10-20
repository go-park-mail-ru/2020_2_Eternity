# Create pinterest user with password and grant all privileges
.PHONY: dbinit
dbinit:
	sudo -u postgres psql -f configs/sql/init.sql;


# Create all tables
.PHONY: dbsetup
dbsetup:
	PGPASSWORD='662f2710-4e08-4be7-a278-a53ae86ba7f6' psql -U pinterest_user -h 127.0.0.1 -d pinterest_db -f configs/sql/base.sql -w


# Drop all created tables
.PHONY: dbclear
dbclear:
	echo  "select 'drop table if exists \"' || tablename || '\" cascade;' from pg_tables where schemaname = 'public';" > configs/sql/1.sql;
	PGPASSWORD='662f2710-4e08-4be7-a278-a53ae86ba7f6' psql -U pinterest_user -h 127.0.0.1 -d pinterest_db -f configs/sql/1.sql -w | grep drop > configs/sql/2.sql;
	PGPASSWORD='662f2710-4e08-4be7-a278-a53ae86ba7f6' psql -U pinterest_user -h 127.0.0.1 -d pinterest_db -f configs/sql/2.sql -w;
	rm configs/sql/1.sql configs/sql/2.sql;


# Build target api
.PHONY: build
build:
	go build -o build/bin/api ./cmd/api/main.go
