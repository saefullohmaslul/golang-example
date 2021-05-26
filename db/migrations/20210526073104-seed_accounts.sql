
-- +migrate Up
INSERT INTO accounts (
    account_number,
    customer_number,
    balance
)
VALUES
    (
        555001,
        1001,
        10000
    ),
    (
        555002,
        1002,
        15000
    );

-- +migrate Down
DELETE FROM accounts;