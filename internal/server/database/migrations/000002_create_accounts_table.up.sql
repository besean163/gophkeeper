CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    name VARCHAR(100),
    login VARCHAR(20),
    password VARCHAR(100),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);