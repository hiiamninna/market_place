BEGIN;

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(128) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(128) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    price INT,
    image_url VARCHAR(255) NOT NULL,
    stock INT,
    condition VARCHAR(10),
    is_purchaseable BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- NOTES
-- ADD owner of the products

COMMIT;