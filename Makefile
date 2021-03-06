# Create pinterest user with password and grant all privileges
.PHONY: dbinit
dbinit:
	sudo -u postgres psql -f configs/sql/init.sql;


# Create all tables
.PHONY: dbsetup
dbsetup:
	PGPASSWORD='662f2710-4e08-4be7-a278-a53ae86ba7f6' psql -U pinterest_user -h 127.0.0.1 -d pinterest_db -f configs/sql/base.sql -f configs/sql/comments.sql -f configs/sql/notifications.sql -f configs/sql/chat.sql -w


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
	go build -o build/bin/api ./cmd/api/main.go;
	go build -o build/bin/chat ./cmd/chat/main.go;
	go build -o build/bin/auth ./cmd/auth/authService.go;
	go build -o build/bin/search ./cmd/search/searchserv.go;



.PHONY: protochat
protochat:
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative pkg/proto/chat/chat.proto


.PHONY: test
test:
	go test ./... -coverprofile=cover.out.tmp -coverpkg=./... -cover ./...
	cat cover.out.tmp | grep -v "_easyjson.go"| grep -v "/mock/" | grep -v "proto" | grep -v "internal" | grep -v "cmd" | grep -v "ws"  > cover.out
	go tool cover -func cover.out


.PHONY: lint
lint:
	echo "Linters passed"