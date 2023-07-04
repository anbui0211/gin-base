# gin-server

### Swaggger

```shell
# Swagger app after run app
http://localhost:8001/swagger/index.html
```

### Docker

```shell
# Run server in docker
docker compose up -d

# Exec container server
docker exec -it backend sh

# Check log server realtime
docker logs -f backend
```

### Run app local

```shell
# Run server
make run-app

# migrate to database: create table and insert rows
make migrations-up

# Reverting the changes made to the previous state
make migrations-down
```

### Document

```shell
# Framework go
https://gin-gonic.com/

# ORM library for Golang
https://gorm.io/

# Swag
https://github.com/swaggo/swag

# Validate
https://github.com/go-ozzo/ozzo-validation

# Migration
https://www.freecodecamp.org/news/database-migration-golang-migrate/

# Xo
https://github.com/xo/xo
```


### Note
```shell
# Xo: Generate JSON struct model from model in database

```


