-- Down migration: add back created_at column to category table
ALTER TABLE category ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
