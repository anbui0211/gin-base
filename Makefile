#!bin/bash
export POSTGRESQL_URL=postgres://anbui:1234@localhost:5432/pgtest?sslmode=disable

make run:
	@go run cmd/myapp/main.go

# make create-migrations name=add_a_column
create-migration:
	@migrate create -ext sql -dir internal/models/pg/migrations -seq $(name)

# make migrations-up num=1
migrations-up:
	@migrate -database ${POSTGRESQL_URL} -path internal/models/pg/migrations up $(num)

# make migrations-down num=1
migrations-down:
	@migrate -database ${POSTGRESQL_URL} -path internal/models/pg/migrations down $(num)



swagger-app:
	swag init -d ./ -g cmd/myapp/main.go \
#    --exclude ./ \
    -o ./docs
