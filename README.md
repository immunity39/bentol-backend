# bentol-backend
backend側の設定等について

## API's
- ログイン +json
POST /login
- ショップ一覧
GET /store
- ショップの弁当一覧
GET /store/{id}
- ショップの弁当選択
GET /store/{id}/bento/{id} or /store/{id}/{id}
- 支払い + json
POST /payment
- 注文キャンセル
DELETE /cancel/{id}

ショップ側の操作は一度後回し
