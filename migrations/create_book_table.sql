create table book(
book_id serial primary key,
author_name varchar(250) not null,
book_title varchar(300) unique not null,
isbn varchar(20) unique not null
);