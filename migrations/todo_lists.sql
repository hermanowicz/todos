create table todo_lists(name varchar(50) not null, user varchar(50));
create index idx_users_list on todo_lists(user);
