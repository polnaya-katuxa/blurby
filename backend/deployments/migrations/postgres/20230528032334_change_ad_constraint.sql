-- +goose Up
-- +goose StatementBegin
alter table ads drop constraint ads_creation_time_check;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table ads
    add constraint ads_creation_time_check
        check (creation_time < current_timestamp::timestamp);
-- +goose StatementEnd
