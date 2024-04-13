.PHONY: run
run: build
	./bin/urlStash

.PHONY: build
build:
	go build -o ./bin/urlStash ./cmd/urlStash/main.go

.PHONY: watch
watch:
	@${HOME}/go/bin/air

.PHONY: test
test:
	go test -v ./...

# DB Operations
.PHONY: clearDB
clearDB:
	@echo "Clearing DB"
	@psql -h localhost -p 5431 urlStash -f ./migrations/init_down.sql

.PHONY: initDB
initDB:
	@echo "Creating DB Tables and References"
	@psql -h localhost -p 5431 urlStash -f ./migrations/init_up.sql

.PHONY: refreshDB
refreshDB: clearDB initDB

.PHONY: refreshAndPopulate
refreshAndPopulate: refreshDB
	@echo "Populating DB with test data"
	@psql -h localhost -p 5431 urlStash -f ./migrations/populate_testdata.sql
