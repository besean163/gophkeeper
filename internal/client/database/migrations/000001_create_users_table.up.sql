CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    login VARCHAR(20) UNIQUE,
    password VARCHAR(100),
    token VARCHAR(100),
    created_at INTEGER
);