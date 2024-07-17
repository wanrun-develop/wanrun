CREATE TABLE IF NOT EXISTS tag_mst (
    tag_id serial primary key,
    tag_name varchar(64) not null,
    description text
);
