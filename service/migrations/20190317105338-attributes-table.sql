-- +migrate Up
CREATE TABLE attributes(
             attribute_id UUID NOT NULL PRIMARY KEY,
             version UUID NOT NULL,
             created TIMESTAMP NOT NULL,
             updated TIMESTAMP NOT NULL,
             name TEXT NOT NULL,
             description TEXT NOT NULL
);

-- +migrate Down
DROP TABLE attributes;

