-- +goose Up
create table session (
	id serial primary key,
	user_id int not null references "user"(id)
		on delete cascade
);

-- +goose Down
drop table if exists session;
