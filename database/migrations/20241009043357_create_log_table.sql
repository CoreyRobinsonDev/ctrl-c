-- +goose Up
create table log (
	id serial primary key,
	type char(7), 
	message varchar,
	action varchar,
	created_at timestamp default now()
);

-- +goose Down
drop table if exists log;
