-- +goose Up
-- +goose StatementBegin
alter table transactions add column type varchar(10) not null
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table transactions drop column type
-- +goose StatementEnd
