create table users (
  user_id integer primary key,
  user_name varchar(100),
  email_address varchar(100),
  user_password varchar(50)
);

create unique index uk_email_address on users(email_address);

create table item_category (
    category_id integer primary key,
    category_name varchar(100)
);

create table items (
    item_id integer primary key,
    item_name varchar(100),
    category_id integer not null,
    item_owner_id integer not null,
    constraint fk_category_id
        foreign key (category_id) references item_category(category_id),
    constraint fk_item_owner_id
        foreign key (item_owner_id) references users(user_id)
);

