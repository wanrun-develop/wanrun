CREATE TYPE grant_type_enum AS ENUM ('PASSWORD', 'OAUTH');

CREATE TABLE IF NOT EXISTS auth_dog_owners (
    auth_dog_owner_id serial primary key,
    dog_owner_id bigint not null,
    password varchar(256),  -- パスワード認証用のパスワード。OAuth認証の場合はNULL。
    email varchar(255),  -- パスワード認証で使うemail。OAuth認証の場合はNULL。
    phone_number varchar(15),  -- パスワード認証で使うphone_number。OAuth認証の場合はNULL。
    grant_type grant_type_enum not null,  -- 認証方式のEnum型 (PASSWORD または OAUTH)
    login_at timestamp,

   -- grant_typeが'PASSWORD'の場合にのみemailかphone_numberのどちらかが必須
    CHECK ((grant_type = 'PASSWORD' AND (email IS NOT NULL OR phone_number IS NOT NULL)) OR grant_type = 'OAUTH')
);
