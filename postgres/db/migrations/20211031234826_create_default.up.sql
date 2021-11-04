CREATE TABLE users(
    id serial not null unique,
    Phone VARCHAR(11) UNIQUE NOT NULL
);

CREATE TABLE code_users(
    Users_Id INTEGER REFERENCES users (Id),
    Auth_Code VARCHAR(4),
    FOREIGN KEY (Users_Id) REFERENCES users (Id)  ON DELETE CASCADE
);

CREATE TABLE products(
    id serial not null unique,
    name varchar(255) UNIQUE NOT NULL
);
