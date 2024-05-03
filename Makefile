include .env
SERVER_TARGET='build/mf_server'

MIGRATION_TARGET=db/migrations
DATABASE_URL=${DB_URL}
MIGRATOR=migrate -verbose -path ${MIGRATION_TARGET} -database ${DATABASE_URL}

all:
	go run cmd/server/main.go

build:
	go build cmd/server/main.go
run:
	./main
script:
	go run cmd/script/main.go



gen_migration:
	migrate -verbose create -ext sql -tz utc -dir ${MIGRATION_TARGET} ${name} 
migrate_up:
	${MIGRATOR} up
migrate_up_to:
	${MIGRATOR} up ${version}
migrate_down:
	${MIGRATOR} down
migrate_down_to:
	${MIGRATOR} down ${version}
migrate_goto:
	${MIGRATOR} goto ${version}
migrate_drop:
	${MIGRATOR} drop
migrate_force:
	${MIGRATOR} force ${version}
migrate_version:
	${MIGRATOR} version


# these needs update in case of new project
generate_models:
	### Extra options for tables-to-go
  # host of database (default "127.0.0.1") -h string
	# password of database -p string
	# user of database -u string
	# port of database -port
	mkdir -p app/models/new_gen
	tables-to-go -v -t mysql -d aero_development  -u root -tags-no-db -pn models -fn-format s -of app/models/new_gen
	gomodifytags -all  -add-tags json -w -dir app/models/new_gen
	# copy the required model in the model directory and then delete new_gen

schema_dump:
	# Options that can be provided
	# mysqldump -h yourhostnameorIP -u root -p --no-data dbname > schema.sql
	mysqldump -u root --no-data --dump-date aero_development > db/schema.sql