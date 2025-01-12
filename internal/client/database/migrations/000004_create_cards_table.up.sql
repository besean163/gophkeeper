CREATE TABLE cards (
    uuid UUID PRIMARY KEY,
    user_id INTEGER,
    name VARCHAR(100),
    number INTEGER,
    exp VARCHAR(10),
    cvv INTEGER,
    created_at INTEGER,
    updated_at INTEGER,
    deleted_at INTEGER,
    synced_at INTEGER
);