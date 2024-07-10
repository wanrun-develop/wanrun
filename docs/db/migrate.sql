\c dog_runner;

create schema dev;

create role developer login SUPERUSER password 'developer';

set search_path = 'dev';

CREATE TABLE IF NOT EXISTS dog_owners (
    dog_owner_id serial primary key,
    name varchar(128) not null,
    email varchar(255) not null,
    image text,
    sex char(1),
    reg_at timestamp not null,
    upd_at timestamp not null
);

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

CREATE TABLE IF NOT EXISTS dog_type_mst (
    dog_type_id serial primary key,
    name varchar(64) not null
);

CREATE TABLE IF NOT EXISTS injection_certifications (
    injection_certification_id serial primary key,
    dog_id bigint not null,
    type int not null,
    file text not null,
    reg_at timestamp not null,
    upd_at timestamp not null
);

CREATE TABLE IF NOT EXISTS dogrun_managers (
    dogrun_manager_id serial primary key,
    name varchar(128),
    email varchar(255) not null,
    reg_at timestamp not null,
    upd_at timestamp not null
);

CREATE TABLE IF NOT EXISTS dogruns (
    dogrun_id serial primary key,
    dogrun_manager_id bigint,
    name varchar(256) not null,
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

CREATE TABLE IF NOT EXISTS dogrun_images (
    dogrun_image_id serial primary key,
    dogrun_id bigint not null,
    image text not null,
    "order" int,
    upload_at timestamp
);

CREATE TABLE IF NOT EXISTS dogrun_tags (
    dogrun_tag_id serial primary key,
    dogrun_id bigint not null,
    tag int not null
);

CREATE TABLE IF NOT EXISTS tag_mst (
    tag_id serial primary key,
    tag_name varchar(64),
    description text
);


CREATE TABLE IF NOT EXISTS auth_dog_owners (
    auth_dog_owner_id serial primary key,
    dog_owner_id bigint not null,
    password varchar(256),
    grant_type int not null,
    login_at timestamp
);


CREATE TABLE IF NOT EXISTS auth_dogrun_managers (
    auth_dogrun_manager_id serial primary key,
    dogrun_manager_id bigint not null,
    password varchar(256),
    grant_type int not null,
    login_at timestamp
);
