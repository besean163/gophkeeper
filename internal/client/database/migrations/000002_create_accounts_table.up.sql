CREATE TABLE accounts (
    id INTEGER PRIMARY KEY,
    uuid UUID UNIQUE,
    user_id INTEGER,
    name VARCHAR(100),
    login VARCHAR(20),
    password VARCHAR(100),
    created_at INTEGER,
    updated_at INTEGER,
    deleted_at INTEGER,
    synced_at INTEGER
);