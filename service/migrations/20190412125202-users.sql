-- +migrate Up
CREATE TABLE users(
             user_id UUID NOT NULL PRIMARY KEY,
             version UUID NOT NULL,
             created TIMESTAMP NOT NULL,
             updated TIMESTAMP NOT NULL,
             name TEXT NOT NULL,
             email TEXT NULL UNIQUE,
             authentication JSONB NOT NULL
);

-- +migrate Down
DROP TABLE users;
