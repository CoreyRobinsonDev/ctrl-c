-- +goose Up
create table post (
	id serial primary key,
	user_id int references "user"(id),
	description char(256),
	snippet_id int references snippet(id),
	stars int default 0,
	discussions int default 0,
	forks int default 0,
	clones int default 0,
	fork_id int references post(id),
	branch_id int references post(id),
	created_at timestamp default now()
);

-- +goose Down
drop table if exists post;
