CREATE TABLE users(
    id serial not null unique,
    Phone VARCHAR(11) UNIQUE NOT NULL
);

CREATE TABLE codeUsers(
    Id serial not null unique,
    UsersId INTEGER REFERENCES users (Id),
    AuthCode VARCHAR(4),
    FOREIGN KEY (UsersId) REFERENCES users (Id)  ON DELETE CASCADE
);

CREATE TABLE products(
    id serial not null unique,
    name varchar(255) UNIQUE NOT NULL
);
