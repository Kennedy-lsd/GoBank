-- Revert `id` column in `accounts` table to SERIAL/INTEGER
ALTER TABLE accounts ALTER COLUMN id TYPE INTEGER;

-- Revert `id` and `account_id` columns in `balances` table to SERIAL/INTEGER
ALTER TABLE balances ALTER COLUMN id TYPE INTEGER;
ALTER TABLE balances ALTER COLUMN account_id TYPE INTEGER;
