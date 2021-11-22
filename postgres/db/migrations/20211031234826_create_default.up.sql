CREATE TABLE users
(
    id    serial             not null unique,
    Phone VARCHAR(11) UNIQUE NOT NULL
);

CREATE TABLE code_users
(
    id        serial             not null unique,
    Phone     VARCHAR(11) UNIQUE NOT NULL,
    Auth_Code VARCHAR(4)
);

CREATE TABLE products
(
    id   serial              not null unique,
    name varchar(255) UNIQUE NOT NULL
);
