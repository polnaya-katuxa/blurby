create role admin login password 'password1';
create role targetologist login password 'password2';
create role client login password 'password3';

grant insert on clients to client;
grant insert on events to client;

grant insert, select on users to targetologist;
grant insert, select on ads to targetologist;
grant insert, select on event_types to targetologist;
grant select on events to targetologist;
grant insert, select, update on schedules to targetologist;
grant select on clients to targetologist;

grant insert, select, delete, update on users to admin;
grant select, delete on clients to admin;
