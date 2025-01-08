CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    uuid UUID UNIQUE,
    user_id INTEGER,
    name VARCHAR(100),
    content TEXT,
    created_at INTEGER,
    updated_at INTEGER
);