-- +goose Up
-- +goose StatementBegin
create extension "uuid-ossp";

create table users (
                       uuid uuid default uuid_generate_v4() primary key,
                       login text unique not null,
                       password text not null,
                       is_admin bool default false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
drop extension if exists "uuid-ossp";
-- +goose StatementEnd
