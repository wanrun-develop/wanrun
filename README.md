# DogRunner

## docker command

### Postgres立ち上げとEchoの立ち上げ
docker compose up -d --build

## Local db migration

### FYI
https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md

### SQLファイル作成コマンド
```
*絶対パスで指定ができなかったので注意(ルートフォルダからやって)
migrate create -ext sql -dir migrate/migration_sql -seq {create_table_name}
```

### DB作成コマンド
dbディレクトリの中に移動して下記のコマンドを実行
`MIGRATION_DIRECTION`の環境変数で識別をしてる。

**DB作成**
```
MIGRATION_DIRECTION=up go run migrate.go
```

**DB削除**
```
MIGRATION_DIRECTION=down go run migrate.go
```
