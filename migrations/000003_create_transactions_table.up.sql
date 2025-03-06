CREATE TABLE transactions
(
    id               SERIAL PRIMARY KEY,
    user_id          INT REFERENCES users (id) ON DELETE CASCADE,
    amount           DECIMAL(10, 2) NOT NULL,
    transaction_type VARCHAR(10) CHECK (transaction_type IN ('credit', 'debit')),
    category_id      INT            REFERENCES categories (id) ON DELETE SET NULL,
    timestamp        TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_transactions_user_id ON transactions (user_id);
CREATE INDEX idx_transactions_timestamp ON transactions (timestamp);
