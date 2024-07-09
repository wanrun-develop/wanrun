# DogRunner

## docker command

### Postgres立ち上げとEchoの立ち上げ
docker compose up -d --build

## Local db migration

```
絶対パスで指定ができなかったので注意
migrate create -ext sql -dir .db/migrations -seq {create_table_name}
```
