-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, email, password_hash, created_at, updated_at) VALUES
    ('1b5d0862-0a8e-4ca2-8dac-5437e2086bf9', 'john.doe@example.com', '$2a$12$4ebqCzeiKVOWEM3z.7Wv6u7UN2sItkolsNdSwohbYZS4P3WOaVpmm', NOW(), NOW()),
    ('f0d07f6f-2e3a-4621-8c3c-1d98576dfcb3', 'jane.smith@example.com', '$2a$12$4ebqCzeiKVOWEM3z.7Wv6u7UN2sItkolsNdSwohbYZS4P3WOaVpmm', NOW(), NOW()),
    ('c2c68525-8c5e-4421-a89d-6f0bb32e7738', 'bob.wilson@example.com', '$2a$12$4ebqCzeiKVOWEM3z.7Wv6u7UN2sItkolsNdSwohbYZS4P3WOaVpmm', NOW(), NOW());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users 
WHERE id IN (
    '1b5d0862-0a8e-4ca2-8dac-5437e2086bf9',
    'f0d07f6f-2e3a-4621-8c3c-1d98576dfcb3',
    'c2c68525-8c5e-4421-a89d-6f0bb32e7738'
);
-- +goose StatementEnd
