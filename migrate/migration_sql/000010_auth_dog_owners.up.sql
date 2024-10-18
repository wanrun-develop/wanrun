-- CREATE TYPE IF NOT EXISTSのオプションがないため
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'grant_type_enum') THEN
        CREATE TYPE grant_type_enum AS ENUM ('PASSWORD', 'OAUTH');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS auth_dog_owners (
    auth_dog_owner_id serial primary key,         -- PK
    dog_owner_id bigint not null,                 -- dog_ownersへの外部キー
    grant_type grant_type_enum not null,          -- 認証方式のEnum型 (PASSWORD または OAUTH)
    access_token VARCHAR(512),                    -- アクセストークン（OAuth 認証の場合のみ）
    refresh_token VARCHAR(512),                   -- リフレッシュトークン（OAuth 認証の場合のみ）
    access_token_expiration timestamp,            -- アクセストークンの有効期限
    refresh_token_expiration timestamp,           -- リフレッシュトークンの有効期限（オプション）
    login_at timestamp
);
