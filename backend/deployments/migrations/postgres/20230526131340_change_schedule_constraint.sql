-- +goose Up
-- +goose StatementBegin
alter table schedules drop constraint schedules_next_time_check;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table schedules
add constraint schedules_next_time_check
check ( next_time > current_timestamp::timestamp );
-- +goose StatementEnd
