insert into users
(
user_name, user_surname, user_password
)values
(
'uluk', 'toychuev', 'uluk_1424'
);


insert into book 
(
author_name,
book_title,
isbn
)
values
(
'Ray Bradbury',
'Fahrenheit 451',
'978-1451673319'
);



insert into users_book 
(
user_id, 
book_id
)
values
(
(select users.id from users where users.user_name = 'uluk'), 
(select book.id from book where book.book_title = 'Fahrenheit 451')
);