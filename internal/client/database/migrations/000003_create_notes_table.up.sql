CREATE TABLE notes (
    uuid UUID PRIMARY KEY,
    user_id INTEGER,
    name VARCHAR(100),
    content TEXT,
    created_at INTEGER,
    updated_at INTEGER,
    deleted_at INTEGER,
    synced_at INTEGER
);