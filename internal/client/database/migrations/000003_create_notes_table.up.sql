CREATE TABLE notes (
    id INTEGER PRIMARY KEY,
    uuid UUID UNIQUE,
    user_id INTEGER,
    name VARCHAR(100),
    content TEXT,
    created_at INTEGER,
    updated_at INTEGER,
    deleted_at INTEGER,
    synced_at INTEGER
);