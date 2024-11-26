-- Update `id` column in `accounts` table
ALTER TABLE accounts ALTER COLUMN id TYPE BIGINT;

-- Update `id` and `account_id` columns in `balances` table
ALTER TABLE balances ALTER COLUMN id TYPE BIGINT;
ALTER TABLE balances ALTER COLUMN account_id TYPE BIGINT;
