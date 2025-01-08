CREATE TABLE cards (
    id SERIAL PRIMARY KEY,
    uuid UUID UNIQUE,
    user_id INTEGER,
    name VARCHAR(100),
    number INTEGER,
    exp VARCHAR(10),
    cvv INTEGER,
    created_at INTEGER,
    updated_at INTEGER
);