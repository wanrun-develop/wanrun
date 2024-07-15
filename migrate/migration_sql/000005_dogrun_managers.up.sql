CREATE TABLE IF NOT EXISTS dogrun_managers (
    dogrun_manager_id serial primary key,
    name varchar(128),
    email varchar(255) not null,
    reg_at timestamp not null,
    upd_at timestamp not null
);
