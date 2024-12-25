CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    login VARCHAR(20),
    password VARCHAR(100),
    token VARCHAR(100),
    create_at TIMESTAMP
);