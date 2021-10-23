BINARY=./bin/service

build:
	@echo "+ $@"
	go build -mod vendor -v -o ${BINARY}

run:
	@echo "+ $@"
	${BINARY}

clean:
	@echo "+ $@"
	rm -f bin/*

createdb:
	# PGHOST=localhost PGUSER=postgres PGPORT=5432 createdb notes

migup:
	sql-migrate up

migdown:
	sql-migrate down