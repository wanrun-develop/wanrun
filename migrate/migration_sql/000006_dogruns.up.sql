CREATE TABLE IF NOT EXISTS dogruns (
    dogrun_id serial primary key,
    place_id varchar(256) not null,
    dogrun_manager_id bigint,
    name varchar(256),
    address varchar(256),
    postcode varchar(8),
    business_day int,
    holiday int,
    open_time time,
    close_time time,
    description text,
    reg_at timestamp not null,
    upd_at timestamp not null
);
