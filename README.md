# OnFES (Online Festival Booth)

サポーターズが開催する、[【技育CAMP2024】ハッカソン Vol.14(テーマあり)【オンライン開催】](https://talent.supporterz.jp/events/6f769c37-5709-4ddc-804e-4c4ff38d2112/)での成果物のひとつ。テーマは「祭」。

本リポジトリは、技育祭などのオンラインフェスのサブブースとして機能する自己紹介、個人開発物とそのリンクを掲載し、ブースを出すことでエンジニアとつながれるWebアプリの、バックエンドのリポジトリである。

**発表スライドリンク**: [KatLab_Online Festival Booth(OnFES) - Google スライド](https://docs.google.com/presentation/d/1iMzPGfuHGFE-4xP5HDIdzmaOzY39G8cRd-i6tG31mSw/edit#slide=id.p)

**フロントエンドのリポジトリ**: [Frontend-Festival-Booth](https://github.com/YKhm20020/Frontend-Festival-Booth)

Webアプリの機能や使い方などに関しては、上記のリポジトリのREADMEをご参考に。



# Backend-Festival-Booth

## 環境構築

1. Dockerをインストール
2. `git clone https://github.com/aridome222/Backend-Festival-Booth.git`
3. `docker network create fes-booth-network`
4. `docker network create fes-booth-pg-network`
5. `docker compose build`

## 開発の流れ

docker コンテナを起動

```
docker compose up
```

docker コンテナを閉じる

```
docker compose down
```

APIサーバは `localhost:8080` , PgAdminは `localhost:81` で立ち上がる

## 技術スタック
- Go
- Gin
- Gorm
- Air
- PostgreSQL
- PgAdmin
- Docker
