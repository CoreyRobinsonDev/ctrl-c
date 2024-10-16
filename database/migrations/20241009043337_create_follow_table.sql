-- +goose Up
create table follow (
	follower_id int not null references "user"(id)
		on delete cascade,
	followee_id int not null references "user"(id)
		on delete cascade
);

-- +goose Down
drop table if exists follow;
