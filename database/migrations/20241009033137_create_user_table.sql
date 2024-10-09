-- +goose Up
create table "user" (
	id serial primary key,
	name char(32) not null,
	handle char(32) not null,
	followers int default 0,
	following int default 0,
	description char(256),
	created_at timestamp default now()
);

-- +goose Down
drop table if exists "user";
