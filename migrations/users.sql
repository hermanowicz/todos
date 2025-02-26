create table users(mail varchar(100), active integer, login_token varchar(50), mail_veryfication_token varchar(50));

create index idx_user_mail on users(mail);
