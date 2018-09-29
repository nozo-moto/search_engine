
-- +migrate Up
ALTER TABLE Page ADD UNIQUE (URL);
-- +migrate Down
ALTER TABLE Page DROP INDEX URL;
