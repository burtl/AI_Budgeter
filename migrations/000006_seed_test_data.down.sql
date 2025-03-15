-- Rollback for seed data
-- Just remove any rows we inserted for the "test" user and associated records

DELETE FROM ai_predictions
WHERE user_id = 1;

DELETE FROM transactions
WHERE user_id = 1;

DELETE FROM budgets
WHERE user_id = 1;

DELETE FROM users
WHERE id = 1;
