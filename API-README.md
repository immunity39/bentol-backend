# 弁当予約販売システム API 仕様

## エンドポイント一覧

### ユーザー登録
- **URL**: `/registration`
- **メソッド**: `POST`
- **リクエスト**
    ```json
    {
        "name": "user-name",
        "password": "user-password",
        "mail": "user-mail"
    }
    ```
- **レスポンス**
    - 成功時: ユーザー登録成功メッセージ
    - 失敗時: エラーメッセージ

### ログイン
- **URL**: `/login`
- **メソッド**: `POST`
- **リクエスト**
    ```json
    {
        "name": "user-name",
        "password": "user-password"
    }
    ```
- **レスポンス**
    - 成功時: ログイン成功メッセージ
    - 失敗時: エラーメッセージ

### 店舗一覧取得
- **URL**: `/store`
- **メソッド**: `GET`
- **リクエスト**: なし
- **レスポンス**
    ```json
    {
        "stores": [
            {"id": 1, "name": "store_1"},
            {"id": 2, "name": "store_2"},
            ...
        ]
    }
    ```

### 店舗の弁当一覧取得
- **URL**: `/store/:id`
- **メソッド**: `GET`
- **リクエスト**: なし
- **レスポンス**
    ```json
    {
        "store": {
            "id": 1,
            "name": "store_1",
            "menues": [
                {"id": 1, "name": "bentou_1", "price": 500, "description": "desc", "is_sold_out": false},
                {"id": 2, "name": "bentou_2", "price": 600, "description": "desc", "is_sold_out": false},
                ...
            ]
        }
    }
    ```

### 予約ポリシー設定
- **URL**: `/store/policy`
- **メソッド**: `POST`
- **リクエスト**
    ```json
    {
        "store_id": 1,
        "date": "2024-06-14",
        "day_of_week": 5,
        "time_slot_interval": 10,
        "max_reservations_per_slot": 5
    }
    ```
- **レスポンス**
    - 成功時: ポリシー設定成功メッセージ
    - 失敗時: エラーメッセージ

### 予約の確認および実行
- **URL**: `/payment`
- **メソッド**: `POST`
- **リクエスト**
    ```json
    {
        "user_id": 1,
        "store_id": 1,
        "menue_id": 1,
        "time": "12:10",
        "date": "2024-06-14",
        "count": 2
    }
    ```
- **レスポンス**
    - 成功時: 予約成功メッセージ
    - 失敗時: エラーメッセージ（予約リミット超過など）

## 注意事項

- すべてのリクエストとレスポンスはJSON形式で行います。
- `user_id`などのユーザー情報は、セキュリティ上の理由から本番環境ではJWTトークンなどの認証機構を利用することを推奨します。
