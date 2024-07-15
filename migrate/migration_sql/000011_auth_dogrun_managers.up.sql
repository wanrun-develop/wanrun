CREATE TABLE IF NOT EXISTS auth_dogrun_managers (
    auth_dogrun_manager_id serial primary key,
    dogrun_manager_id bigint not null,
    password varchar(256),
    grant_type int not null,
    login_at timestamp
);
