# Backend-Festival-Booth

## 環境構築

1. Go および Docker をインストール
2. `git clone https://github.com/aridome222/Backend-Festival-Booth.git`
3. `docker network create fes-booth-network`
4. `docker compose build`

## 開発の流れ

docker コンテナを起動

```
docker compose up
```

docker コンテナを閉じる

```
docker compose down
```

API サーバは `localhost:8080` , PgAdmin は `localhost:81` で立ち上がる

## 技術スタック

- Go
- Gin
- Gorm
- PostgreSQL
- PgAdmin
- Docker

## フロントエンドのリポジトリ

https://github.com/YKhm20020/Frontend-Festival-Booth
