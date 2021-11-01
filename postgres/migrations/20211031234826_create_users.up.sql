CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    phone VARCHAR(11) UNIQUE NOT NULL,
    auth_code VARCHAR(4),
    departure_time TIMESTAMP
);
