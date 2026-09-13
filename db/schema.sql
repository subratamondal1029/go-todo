CREATE TABLE "schema_migrations" (version varchar(128) primary key);
CREATE TABLE todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    completed BOOLEAN DEFAULT false,
    completed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_completed_created_at
ON todos (completed ASC, created_at ASC);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20260913143538'),
  ('20260913145622');
