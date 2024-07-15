# DogRunner

## docker command

### Postgres立ち上げとEchoの立ち上げ
docker compose up -d --build

---

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

**DBテーブル作成**
```
MIGRATION_DIRECTION=up go run migrate.go
```

**DBテーブル削除**
```
MIGRATION_DIRECTION=down go run migrate.go
```

**DB削除**
※Dirty database versionエラーが発生した場合
マイグレーションがどのバージョンで失敗したかを確認します。これはデータベースのマイグレーション履歴テーブル（通常はschema_migrationsまたはflyway_schema_history）を直接クエリすることで確認
```
MIGRATION_DIRECTION=force MIGRATION_VERSION={対象のバージョン番号} go run migrate.go
```

## go doc

### インストール方法
```
go install golang.org/x/tools/cmd/godoc@latest
```

### 閲覧方法
下記のコマンド後に、`http://localhost:6060`を開く
```
godoc -http=:6060
```
