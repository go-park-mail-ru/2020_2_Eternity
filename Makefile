# note: call scripts from /scripts
.PHONY: dbinit
dbinit:
	sudo -u postgres psql -f configs/sql/init.sql


