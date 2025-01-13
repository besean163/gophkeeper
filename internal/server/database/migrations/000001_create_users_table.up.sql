CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    uuid UUID  UNIQUE,
    login VARCHAR(20) UNIQUE,
    password VARCHAR(100),
    created_at INTEGER
);