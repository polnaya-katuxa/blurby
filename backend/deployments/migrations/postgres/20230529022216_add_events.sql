-- +goose Up
-- +goose StatementBegin
create table events (
     uuid uuid default uuid_generate_v4() primary key,
     time timestamp not null,
     client_id uuid references clients (uuid) on delete set null,
     alias text references event_types (alias) not null check ( length(alias) > 5 )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists events;
-- +goose StatementEnd
