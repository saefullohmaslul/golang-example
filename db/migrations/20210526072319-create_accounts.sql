
-- +migrate Up
CREATE TABLE accounts
(
    account_number INTEGER PRIMARY KEY,
    customer_number INTEGER NOT NULL,
    balance INTEGER DEFAULT 0
);

ALTER TABLE accounts
ADD CONSTRAINT customer_acount_fk
FOREIGN KEY (customer_number) REFERENCES customers (customer_number)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE accounts
DROP CONSTRAINT customer_acount_fk;
DROP TABLE accounts;