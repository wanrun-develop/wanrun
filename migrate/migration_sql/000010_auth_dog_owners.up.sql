CREATE TABLE IF NOT EXISTS auth_dog_owners (
    auth_dog_owner_id serial primary key,
    dog_owner_id bigint not null,
    password varchar(256),
    grant_type int not null,
    login_at timestamp
);
