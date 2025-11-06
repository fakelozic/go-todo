-- +goose Up
CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY,
	task TEXT NOT NULL,
	status BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE tasks;