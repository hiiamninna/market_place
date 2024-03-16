CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    price INT,
    image_url VARCHAR(255) NOT NULL,
    stock INT,
    condition VARCHAR(10),
    is_purchaseable BOOLEAN DEFAULT TRUE,
    tags VARCHAR[],
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS bank_accounts (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    account_name VARCHAR(255) NOT NULL,
    account_number VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    bank_account_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    total_payment INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);