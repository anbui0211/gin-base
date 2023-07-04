create table users (
   id SERIAL PRIMARY KEY,
   user_id VARCHAR(255) NOT NULL,
   name VARCHAR(100) NOT NULL,
   email VARCHAR(100) NOT NULL,
   username VARCHAR(100) NOT NULL,
   password VARCHAR(100) NOT NULL,
   status VARCHAR(100) NOT NULL,
   phone VARCHAR(100) NOT NULL,
   created_at  TIMESTAMP NOT NULL ,
   updated_at TIMESTAMP NOT NULL
)
