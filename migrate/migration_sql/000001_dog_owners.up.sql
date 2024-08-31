CREATE TABLE IF NOT EXISTS dog_owners (
    dog_owner_id serial primary key,
    name varchar(128) not null,
    email varchar(255) unique not null,
    image text,
    sex char(1),
    reg_at timestamp not null,
    upd_at timestamp not null
);
