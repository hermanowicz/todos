create table todos
(user varchar(50), todo_title varchar(100) not null, todo_body text, todo_list varchar(50), status integer default 0);

create index idx_user on todos(user);
