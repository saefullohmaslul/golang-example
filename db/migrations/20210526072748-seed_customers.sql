
-- +migrate Up
INSERT INTO customers (
    customer_number,
    name
)
VALUES
    (
        1001,
        'Bob Martin'
    ),
    (
        1002,
        'Linus Torvalds'
    );

-- +migrate Down
DELETE FROM customers;