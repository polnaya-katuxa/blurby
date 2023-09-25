-- +goose Up
-- +goose StatementBegin
alter table schedules add column finished boolean default false not null;
alter table schedules add column periodic boolean default false not null;

create or replace function get_ads_by_span(start timestamp, finish timestamp) returns table(
   uuid uuid,
   content text,
   filters json,
   user_id uuid,
   schedule_id uuid,
   creation_time timestamp,
   next_time timestamp,
   span text,
   finished boolean,
   periodic boolean
)
as $$
begin
    return query (
        select a.uuid as uuid, a.content, a.filters, a.user_id, a.schedule_id,
               a.creation_time, s.next_time, s.span, s.finished, s.periodic
        from ads a join schedules s on s.uuid = a.schedule_id
        where s.next_time between start and finish and s.finished = false
    );
end;
$$
language plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop function if exists get_ads_by_span;
alter table schedules drop column finished;
alter table schedules drop column periodic;
-- +goose StatementEnd
