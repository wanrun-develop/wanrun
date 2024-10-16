CREATE TABLE IF NOT EXISTS dog_owners (
    dog_owner_id serial primary key,
    name varchar(128) not null,
    email varchar(255) unique,  -- emailはユニークだが、NULLも許可される
    phone_number varchar(15) unique,  -- phone_numberもNULLを許可する
    image text,
    sex char(1),
    reg_at timestamp not null,
    upd_at timestamp not null
);
