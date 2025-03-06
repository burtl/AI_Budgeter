-- Rollback for Budgets Table
DROP TABLE IF EXISTS budgets;

-- Rollback for Budgets Index
DROP INDEX IF EXISTS idx_budgets_user_id;
