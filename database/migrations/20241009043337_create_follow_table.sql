-- +goose Up
create table follow (
	follower_id int not null references "user"(id),
	followee_id int not null references "user"(id)
);

-- +goose Down
drop table if exists follow;
