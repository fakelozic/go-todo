-- +goose Up
ALTER TABLE tasks
ADD COLUMN user_id UUID NOT NULL,
ADD CONSTRAINT fk_tasks_user_id
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE tasks
DROP CONSTRAINT fk_tasks_user_id,
DROP COLUMN user_id;