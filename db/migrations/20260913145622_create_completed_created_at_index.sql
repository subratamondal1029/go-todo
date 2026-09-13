-- migrate:up
CREATE INDEX IF NOT EXISTS idx_completed_created_at
ON todos (completed ASC, created_at ASC);

-- migrate:down
DROP INDEX IF EXISTS idx_completed_created_at;