CREATE TABLE ai_predictions
(
    id                 SERIAL PRIMARY KEY,
    user_id            INT REFERENCES users (id) ON DELETE CASCADE,
    predicted_category VARCHAR(50),
    predicted_amount   DECIMAL(10, 2),
    prediction_date    TIMESTAMP DEFAULT NOW()
);

