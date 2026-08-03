-- +goose Up
ALTER TABLE users
RENAME COLUMN name TO first_name;

ALTER TABLE users
ADD COLUMN last_name VARCHAR(100);

-- +goose Down
ALTER TABLE users
DROP COLUMN last_name;

ALTER TABLE users
RENAME COLUMN first_name TO name;
