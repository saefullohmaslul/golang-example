
-- +migrate Up
CREATE TABLE customers
(
    customer_number INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE customers;