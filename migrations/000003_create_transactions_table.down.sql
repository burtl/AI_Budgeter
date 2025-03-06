-- Rollback for Transactions Table
DROP TABLE IF EXISTS transactions;

-- Rollback for Transactions Indexes
DROP INDEX IF EXISTS idx_transactions_user_id;
DROP INDEX IF EXISTS idx_transactions_timestamp;