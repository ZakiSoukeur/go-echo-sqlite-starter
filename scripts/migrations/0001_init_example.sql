-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists authors(
  id   INTEGER PRIMARY KEY,
  name text    NOT NULL,
  bio  text
);
INSERT INTO "authors" ("id", "name") VALUES (1, 'gallant_almeida7');
INSERT INTO "authors" ("id", "name") VALUES (2, 'brave_spence8');
INSERT INTO "authors" ("id", "name") VALUES (99999, 'jovial_chaum1');
INSERT INTO "authors" ("id", "name") VALUES (100000, 'goofy_ptolemy0');
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- Add down migration script here
-- Drop Row Level Security policies
DROP TABLE IF EXISTS authors;

-- +goose StatementEnd
