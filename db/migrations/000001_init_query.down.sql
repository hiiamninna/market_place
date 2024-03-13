-- DROP TABLE
/**
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS bank_accounts;
**/

-- ALTER TABLE
/**
ALTER TABLE products ADD COLUMN tags VARCHAR[];
ALTER TABLE products DROP COLUMN test_tags;
**/