-- Up migration: remove created_at column from category table
ALTER TABLE category DROP COLUMN created_at;
