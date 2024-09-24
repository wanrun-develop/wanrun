CREATE TABLE IF NOT EXISTS oauth_providers (
    oauth_id serial primary key,
    auth_dog_owner_id bigint not null references auth_dog_owners(auth_dog_owner_id),  -- auth_dog_ownersへの外部キー
    provider_name varchar(50) not null,  -- OAuthプロバイダ名（例: 'google', 'facebook' など）
    provider_user_id varchar(255) not null,  -- OAuthプロバイダから提供されるユーザーID
    token varchar(512),  -- アクセストークンやリフレッシュトークン（オプション）
    token_expiration timestamp,  -- トークンの有効期限（オプション）
    login_at timestamp  -- 最後のログイン時間
);
