
-- +migrate Up
ALTER TABLE Page ADD  DESCRIPTION TEXT;
-- +migrate Down
ALTER TABLE Page DROP COLUMN DESCRIPTION;
