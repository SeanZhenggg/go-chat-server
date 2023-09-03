create table users (
   id serial primary key,
   account varchar(50) unique not null,
   password varchar(50) not null,
   nickname varchar(50) unique not null,
   create_at timestamp,
   update_at timestamp
)