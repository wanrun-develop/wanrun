CREATE TABLE IF NOT EXISTS dogs (
    dog_id serial primary key,
    dog_owner_id bigint not null,
    name varchar(128) not null,
    dog_type_id int,
    weight int,
    sex char(1),
    image text,
    reg_at timestamp not null,
    upd_at timestamp not null
);
