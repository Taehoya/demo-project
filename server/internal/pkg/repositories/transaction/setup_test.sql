INSERT INTO users
    (id, nickname, email, password)
VALUES
    (1, "test-nickname", "test-email", "test-password");

INSERT INTO trips
    (id, name, user_id, budget, country_id, description, start_date_time, end_date_time, created_at, updated_at, deleted_at)
VALUES
    (1, 'test-name', 1, 1.0000, 1, 'test-description', '2023-11-13 15:04:05', '2023-11-13 15:04:05', '2023-11-13 14:05:27', '2023-11-13 14:05:27', NULL);

INSERT INTO categories
    (id, trip_id, name, colour, icon)
VALUES
    (1, 1, 'name', 'blue', 'blue-icon');

INSERT INTO transactions
    (id, trip_id, user_id, name, amount, `date`, category_id, description, created_at, updated_at)
VALUES
    (1, 1, 1, 'name', 1.01 , CURRENT_TIMESTAMP, 1, 'test-description', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);