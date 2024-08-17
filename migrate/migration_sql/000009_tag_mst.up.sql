CREATE TABLE IF NOT EXISTS tag_mst (
    tag_id serial primary key,
    tag_name varchar(64) not null,
    description text
);

-- マスターデータの登録
insert into tag_master values (1, "屋外", "屋外のドッグランがあります。");
insert into tag_master values (2, "屋内", "屋内のドッグランがあります。");
insert into tag_master values (3, "有料", "ご利用には料金が発生します。");
insert into tag_master values (4, "無料", "料金なしでご利用いただけます。");
insert into tag_master values (5, "予防接種必須", "ご利用には予防接種証明書が必要です。");
insert into tag_master values (6, "予防接種なし", "予防接種証明書がなくてもご利用いただけます。");
insert into tag_master values (7, "駐車場あり", "施設に駐車場がございます。");
insert into tag_master values (8, "無料駐車場あり", "施設に無料駐車場がございます。");
insert into tag_master values (9, "付近に駐車場あり", "施設付近に駐車場（料金発生）がございます。");
insert into tag_master values (10, "大型犬OK", "大型犬もご利用いただけます。");
insert into tag_master values (11, "大型犬NG", "大型犬はご利用いただけません。");
insert into tag_master values (12, "大型犬専用あり", "大型犬専用のドッグランがございます。");
insert into tag_master values (13, "小型犬専用あり", "小型犬専用のドッグランがございます。");
insert into tag_master values (14, "芝生", "地面が芝生のドッグランです。");
insert into tag_master values (15, "海が見える", "施設から海が見えます。");
insert into tag_master values (16, "洗い場あり", "ワンちゃんの洗い場がございます。");
insert into tag_master values (17, "24時間", "24時間ご利用いただけます。");
insert into tag_master values (18, "ショップ併設", "施設にショップが併設しております。");
insert into tag_master values (19, "カフェ併設", "施設にカフェが併設しております。");
insert into tag_master values (20, "屋根付き", "ドッグランに屋根が付いています。");
insert into tag_master values (21, "絶景", "施設からは絶景が見えます。");
insert into tag_master values (22, "パンツ必須", "ご利用時は、パンツを着用が必須です。");
insert into tag_master values (23, "パンツなしでOK", "ご利用時は、パンツを着用しなくてもOKです。");
insert into tag_master values (24, "高速道路SA・PA", "高速道路のSA/PA内にドッグランがございます。");
insert into tag_master values (25, "道の駅", "道の駅内にドッグランがございます。");
insert into tag_master values (26, "アジリティあり", "アジリティ用ドッグランがございます。");




-- 初期データを考慮して、シーケンスの初期値を設定
ALTER SEQUENCE tag_mst_tag_id_seq RESTART WITH 1000;


