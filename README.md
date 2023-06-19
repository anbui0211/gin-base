# server-warehouse

- Install swag
```shell
# https://github.com/swaggo/echo-swagger
go install github.com/swaggo/swag/cmd/swag@latest

# document
https://github.com/swaggo/swag#declarative-comments-format

# Swagger app after run app
http://localhost:8001/swagger/index.html

```

- Docker 
```shell
# Initialize database and server 
docker compose up -d 
```

- Migrate 
```shell
# Create table and insert rows in database
make migrations-up
```

- Run server 
```shell
make run-app
```

