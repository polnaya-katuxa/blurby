-- +goose Up
-- +goose StatementBegin
create or replace function get_ads_by_span(finish timestamp) returns table(
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
        where s.next_time <= finish and s.finished = false
    );
end;
$$
    language plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
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
