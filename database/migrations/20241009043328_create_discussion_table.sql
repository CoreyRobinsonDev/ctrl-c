-- +goose Up
create table discussion (
	id serial primary key,
	post_id int not null references post(id),
	user_id int references "user"(id),
	reply_id int references discussion(id),
	description char(256) not null,
	created_at timestamp default now()
);

-- +goose Down
drop table if exists discussion;
