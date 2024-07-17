CREATE TABLE IF NOT EXISTS dogrun_tags (
    dogrun_tag_id serial primary key,
    dogrun_id bigint not null,
    tag_id int not null
);
