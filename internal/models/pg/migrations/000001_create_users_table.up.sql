CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table users (
   id uuid default uuid_generate_v4() ,
   name varchar(100),
   email varchar(100),
   status varchar(50),
   created_at timestamp,
   updated_at timestamp,
   deletedAt timestamp


)