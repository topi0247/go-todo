
-- +migrate Up
create table todos (
  id serial primary key,
  title text not null,
  description text not null,
  completed boolean not null default false,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp,
  user_id integer not null references users(id)
);

-- +migrate Down
drop table todos;
