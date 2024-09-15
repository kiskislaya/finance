-- +goose Up
-- +goose StatementBegin
create table if not exists transactions (
    id bigint generated always as identity primary key,
    name text not null,
    amount bigint not null,
    created_at timestamp default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table transactions cascade;
-- +goose StatementEnd
