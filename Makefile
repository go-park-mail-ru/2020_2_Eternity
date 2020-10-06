.PHONY: dbsetup
dbsetup:
	sudo -u postgres psql -f configs/sql/init.sql;
	export PGPASSWORD='662f2710-4e08-4be7-a278-a53ae86ba7f6';
	psql -U pinterest_user -h 127.0.0.1 -d pinterest_db -f configs/sql/base.sql -w



