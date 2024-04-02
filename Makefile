run: build
	./bin/urlStash
build:
	go build -o bin/
watch:
	@${HOME}/go/bin/air
test:
	go test -v ./...

# DB Operations
clearDB:
	@echo "Clearing DB"
	@psql -h localhost -p 5431 urlStash -f ./migrations/init_down.sql
initDB:
	@echo "Creating DB Tables and References"
	@psql -h localhost -p 5431 urlStash -f ./migrations/init_up.sql
refreshDB: clearDB initDB
refreshAndPopulate: refreshDB
	@echo "Populating DB with test data"
	@psql -h localhost -p 5431 urlStash -f ./migrations/populate_testdata.sql
