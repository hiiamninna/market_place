BEGIN;

CREATE TABLE users (
    id          SERIAL PRIMARY KEY NOT NULL,
    name        VARCHAR(255) NOT NULL,
    username    VARCHAR(255) NOT NULL,
    password    VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP,
    deleted_at  TIMESTAMP
);

COMMIT;