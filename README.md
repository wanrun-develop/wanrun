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
migrate create -ext sql -dir db/migrations -seq {create_table_name}
```
