#!bin/bash
export POSTGRESQL_URL=postgres://anbui:1234@localhost:5432/pgtest?sslmode=disable

run-app:
	@go run cmd/myapp/main.go

# make create-migrations name=add_a_column
create-migration:
	@migrate create -ext sql -dir db/migrations -seq $(name)

# make migrations-up num=1
migrations-up:
	@migrate -database ${POSTGRESQL_URL} -path db/migrations up $(num)

# make migrations-down num=1
migrations-down:
	@migrate -database ${POSTGRESQL_URL} -path db/migrations down $(num)

# generate enity file from database
xo:
	mkdir -p src/model/xo
	rm -f src/model/xo/*
	xo schema postgres://anbui:1234@localhost:5432/pgtest?sslmode=disable --out=src/model/xo

swagger-app:
	swag init -d ./ -g cmd/myapp/main.go \
#    --exclude ./ \
    -o ./docs

