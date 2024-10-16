-- +goose Up
create table clone (
	id serial primary key,
	cloned_by_id int not null references "user"(id)
		on delete cascade,
	cloned_from_id int not null references post(id)
);

-- +goose Down
drop table if exists clone;
