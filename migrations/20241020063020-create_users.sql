
-- +migrate Up
create table if not exists users (
  id serial primary key,
  uuid varchar(255) not null unique,
  name varchar(255) not null,
  email varchar(255) not null unique,
  password_digest varchar(255) not null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp
);

-- +migrate Down
drop table users;
