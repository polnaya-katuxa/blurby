-- +goose Up
-- +goose StatementBegin
create type gender as enum (
    'female',
    'male'
    );

create table clients (
                         uuid uuid default uuid_generate_v4() primary key,
                         name text not null,
                         surname text not null,
                         patronymic text,
                         gender gender not null,
                         birth_date timestamp not null check (birth_date < current_timestamp::timestamp and birth_date > '01.01.1900 00:00:00'::timestamp),
                         registration_date timestamp not null check (registration_date < current_timestamp::timestamp and registration_date > clients.birth_date),
                         email text unique not null check ( email like '%@%.%' ),
                         data json
);

create table event_types (
    uuid uuid default uuid_generate_v4() primary key,
    name text not null,
    alias text unique not null check ( length(alias) > 5 )
);

create table schedules (
                           uuid uuid default uuid_generate_v4() primary key,
                           next_time timestamp check ( next_time > current_timestamp::timestamp) not null,
                           span text not null
);

create table ads (
    uuid uuid default uuid_generate_v4() primary key,
    content text not null,
    filters json,
    user_id uuid references users (uuid) on delete set null,
    schedule_id uuid unique references schedules (uuid) on delete cascade,
    creation_time timestamp not null check (creation_time < current_timestamp::timestamp)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ads;
drop table if exists schedules;
drop table if exists event_types;
drop table if exists clients;
drop type if exists gender;
-- +goose StatementEnd
