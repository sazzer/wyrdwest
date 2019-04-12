-- +migrate Up
CREATE TABLE clients(
             client_id UUID NOT NULL PRIMARY KEY,
             version UUID NOT NULL,
             created TIMESTAMP NOT NULL,
             updated TIMESTAMP NOT NULL,
             name TEXT NOT NULL,
             owner_id UUID NOT NULL REFERENCES users(user_id) ON DELETE RESTRICT ON UPDATE RESTRICT,
             client_secret UUID NOT NULL,
             redirect_uris TEXT[] NOT NULL,
             response_types TEXT[] NOT NULL,
             grant_types TEXT[] NOT NULL
);

-- +migrate Down
DROP TABLE oauth2_clients;
