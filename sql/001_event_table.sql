create table events (
    id serial primary key,
    description text not null,
    start_time timestamp not null,
    duration interval SECOND(0)
)