CREATE TABLE IF NOT EXISTS dog_owners (
    dog_owner_id serial primary key,    -- PK
    name varchar(128) not null,         -- dog_ownerの名前
    image text,                         -- dog_ownerの写真
    sex char(1),                        -- 性別
    reg_at timestamp not null,          -- 登録日
    upd_at timestamp not null           -- 更新日
);
