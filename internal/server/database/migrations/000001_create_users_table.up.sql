CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(20),
    password VARCHAR(100),
    created_at TIMESTAMP
);