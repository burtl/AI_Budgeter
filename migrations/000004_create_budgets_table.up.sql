CREATE TABLE budgets
(
    id           SERIAL PRIMARY KEY,
    user_id      INT REFERENCES users (id) ON DELETE CASCADE,
    category_id  INT REFERENCES categories (id) ON DELETE CASCADE,
    limit_amount DECIMAL(10, 2) NOT NULL,
    created_at   TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_budgets_user_id ON budgets(user_id);
