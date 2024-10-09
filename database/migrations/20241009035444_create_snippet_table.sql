-- +goose Up
create table snippet (
	id serial primary key,
	user_id int references "user"(id),
	content char(8192) not null,
	language char(32) not null
);

-- +goose Down
drop table if exists snippet;
