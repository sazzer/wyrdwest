CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY,
    version UUID NOT NULL,
    created TIMESTAMP NOT NULL,
    updated TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

