# Diary-Note App

ひとこと日記アプリです。  
提示された質問に対して、日々回答を記録してみましょう！

![記録一覧](docs/screenshots1.png)
![投稿画面](docs/screenshots1.png)
![質問の投稿](docs/screenshots1.png)

# Tech Stack

- Backend
  - Go
  - ORM: GORM
  - DB: Postgres
  - Server: Gin
- Frontend
  - Typescript
  - Framework: Vite React
  - Design Frame: Antdesign

# How to Setup

1. このリポジトリーをクローンする
2. PC 上でこのプロジェクトを開く
3. Go サーバーの起動に必要なライブラリをインストール（# Backend: dependencies を参照）
4. 開発用サーバーの起動

```
$ go run main.go
```

5. フロントエンドの依存関係をインストール

```
$ cd frontend
$ yarn
```

6. 開発用フロントエンドの起動

```
yarn dev
```

## Backend: dependencies

```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/joho/godotenv
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v5
```
