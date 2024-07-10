CREATE TABLE dogs (
    dog_id SERIAL PRIMARY KEY,         -- 犬のID（シリアル）
    owner_id INT REFERENCES owners(owner_id),  -- 飼い主のIDへの外部キー参照
    name VARCHAR(64),                  -- 犬の名前
    dog_type_id INT REFERENCES dog_types(dog_type_id),  -- 犬種のIDへの外部キー参照
    weight INT,                        -- 体重（整数）
    sex INT,                           -- 性別（整数）
    image TEXT,                        -- 写真のURLやパス（テキスト）
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 作成日時（デフォルトは現在時刻）
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- 更新日時（デフォルトは現在時刻）
);
