CREATE TABLE products
(
    id serial not null unique,
    name varchar(255) not null
);

CREATE TABLE users
(
    id serial not null unique,
    phone varchar(11) not null
);
