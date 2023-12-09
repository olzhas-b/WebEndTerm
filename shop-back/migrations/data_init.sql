
insert into category(name)
values ('smartphone');

insert into product(name, category_id, image, price)
values ('android', 1, 'https://iphone.jpg', 1000);


insert into product(name, category_id, image, price)
values ('ipone', 1, 'https://iphone.jpg', 1000);

insert into user_favorite(user_id, product_id)
values (1, 1);

insert into user_cart(user_id, product)
values (1, 1);