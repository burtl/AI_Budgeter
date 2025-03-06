CREATE TABLE categories
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50) UNIQUE NOT NULL,
    description TEXT
);

-- Default categories
INSERT INTO categories (name, description)
VALUES ('Groceries', 'Food shopping'),
       ('Rent', 'Monthly rent payments'),
       ('Entertainment', 'Movies, games, outings');
