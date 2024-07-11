CREATE TABLE IF NOT EXISTS dogrun_images (
    dogrun_image_id serial primary key,
    dogrun_id bigint not null,
    image text not null,
    order int,
    upload_at timestamp
);
