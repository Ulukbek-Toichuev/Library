begin;
create table users
(
id serial primary key,
user_name varchar(250) not null,
user_surname varchar(250) not null,
user_password varchar(250) not null
);


create table book(
id serial primary key,
author_name varchar(250) not null,
book_title varchar(300) unique not null,
isbn varchar(20) unique not null
);

create table users_book
(
id serial primary key,
user_id int references users(id) on delete cascade unique not null,
book_id int references book(id) on delete cascade unique not null
);
rollback;