# bentol-backend
backend側の設定等について

## API's
- ログイン +json
POST /login
- ショップ一覧
GET /shop
- ショップの弁当一覧
GET /shop/{id}
- ショップの弁当選択
GET /shop/{id}/bento/{id} or /shop/{id}/{id}
- 支払い + json
POST /payment
- 注文キャンセル
DELETE /cancel/{id}

ショップ側の操作は一度後回し
