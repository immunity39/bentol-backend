# bentol-backend
backend側の設定等について

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
