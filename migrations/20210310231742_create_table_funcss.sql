-- +goose Up
CREATE TABLE funcss(
    UUID uuid primary key,
    id integer,
    css_hex text,
    name text,
    author text
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE funcss;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
