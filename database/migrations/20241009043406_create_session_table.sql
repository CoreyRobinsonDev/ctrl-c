-- +goose Up
create table session (
	id serial primary key,
	user_id int not null references "user"(id)
);

-- +goose Down
drop table if exists session;
