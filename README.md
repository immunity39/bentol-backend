# Bentol Backend

このプロジェクトは、弁当予約システムのバックエンドサービスです。Go言語とGinウェブフレームワークを使用しています。

## フォルダ構成

- `cmd/app-bentol`: アプリケーションのエントリーポイント。
- `domain/model`: データモデル。
- `domain/repository`: データアクセスロジック。
- `handler`: HTTPハンドラー。
- `handler/validator`: バリデーションロジック。
- `infrastructure`: データベース接続とセットアップ。
- `usecase`: ビジネスロジック。

## セットアップ

1. リポジトリをクローンします。
2. `go mod tidy` を実行して依存関係をインストールします。
3. `infrastructure/db.go` でデータベース接続を設定します。
4. `go run cmd/app-bentol/main.go` を使用してアプリケーションを実行します。


## API's
APIについての細かい設定は ./API-README.mdに記述があります。
- ログイン
POST /login
- ショップ一覧
GET /store
- ショップの弁当一覧
GET /store/{id}
- ショップの弁当選択
GET /menue/{id}
- 支払い
POST /payment
- 注文キャンセル
DELETE /cancel/{id}

ショップ側の操作は一度後回し
