# 弁当予約販売システム API 仕様

## もしhttps://での通信が可能であるならば、http:での通信を除外してもよいかも。

## エンドポイント一覧

## ユーザー関連エンドポイント
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
        "mail": "user@email.com",
        "password": "user-password"
    }
    ```
- **レスポンス**
    - 成功時: ログイン成功メッセージ
    - 失敗時: エラーメッセージ

## ログアウト
ログアウト処理
- **URL** `/logout`
- **メソッド** `POST`
- **リクエスト**
    ```json
    {
        "user_id": 1
    }
    ```
- **レスポンス**
    - 成功時: ログアウト成功
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
- **URL**: `/store/{id}`
- **メソッド**: `GET`
- **リクエスト**: なし
- **レスポンス**
    ```json
    {
        "menues": [
            {
                "ID": 1,
                "StoreID": 1,
                "Name": "bentou-1",
                "Price": 600,
                "Description": "食べれます",
                "CreatedAt": "2024-07-07T21:32:23.713+09:00",
                "UpdatedAt": "2024-07-07T21:34:29.041+09:00"
            },
            {
                "ID": 2,
                "StoreID": 1,
                "Name": "bentou-2",
                "Price": 550,
                "Description": "derisious!",
                "CreatedAt": "2024-07-09T01:49:27.875+09:00",
                "UpdatedAt": "2024-07-09T01:49:27.875+09:00"
            }
        ]
    }
    ```

## 店舗関連エンドポイント
### 販売者登録
- **URL**: `/store/register`
- **メソッド**: `POST`
- **リクエスト**
    ```json
    {
        "store_name": "New Store",
        "mail": "store@email.com",
        "password": "store-password"
    }
    ```
- **レスポンス**
    - 成功時: 販売者登録成功メッセージ
    - 失敗時: エラーメッセージ

### 販売者ログイン
- **URL**: `/store/login`
- **メソッド**: `POST`
- **リクエスト**
    ``` json
    {
        "mail": "store@email.com",
        "password": "store_password"
    }
    ```
- **レスポンス**
    - **成功時**: 販売者ログイン成功メッセージ
    - **失敗時**: エラーメッセージ

### ログアウト
ログアウト処理
- **URL** `/store/logout`
- **メソッド** `POST`
- **リクエスト**
    ```json
    {
    }
    ```
- **レスポンス**
    - 成功時: ログアウト成功
    - 失敗時: エラーメッセージ


### 店舗基本ポリシー設定
- **URL**: `/store/{id}/update`
- **メソッド**: `PUT`
- **リクエスト**
    ```json
    {
        "store_id": 1,
        "policy": {
            "day_of_week": 0,
            "time_slot_interval": 10,
            "max_reservations_per_slot": 5,
            "store_start_time": 11:00,
            "store_end_time": 15:00
        }
    }
    ```
- **レスポンス**
    - 成功時: 店舗情報更新成功メッセージ
    - 失敗時: エラーメッセージ

### 特定予約ポリシー設定
- **URL**: `/store/{id}/policy`
- **メソッド**: `POST`
- **リクエスト**
    ```json
    {
        "store_id": 1,
        "date": "2024-06-14",
        "day_of_week": 5,
        "time_slot_interval": 10,
        "max_reservations_per_slot": 5,
        "store_start_time": 10:00,
        "store_end_time": 18:00
    }
    ```
- **レスポンス**
    - 成功時: ポリシー設定成功メッセージ
    - 失敗時: エラーメッセージ

## メニュー関連エンドポイント
### メニュー追加
- **URL**: `/menue/create`
- **メソッド**: `POST`
- **リクエスト**
    ```json
    {
        "store_id": 1,
        "name": "new bentou",
        "price": 500,
        "description": "str 1",
        "is_sold_out": false
    }
    ```
- **レスポンス**
    - 成功時: メニュー追加成功メッセージ
    - 失敗時: エラーメッセージ
    
### メニュー更新
- **URL**: `/menue/{id}/update`
- **メソッド**: `PUT`
- **リクエスト**
    ```json
    {
        "name": "update bentou",
        "price": 600,
        "descripton": "update",
        "is_sold_out": 0
    }
    ```
- **レスポンス**
    - 成功時: メニュー更新成功メッセージ
    - 失敗時: エラーメッセージ

## 予約関連エンドポイント
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
    - 成功時: paypayの支払先URL
    - 失敗時: エラーメッセージ（予約リミット超過など）

## 予約確認エンドポイント
### 予約確認
## **ポーリングの設定はfront側で行う**
- **URL**: `/store/reservation`
- **メソッド**: `GET`
- **リクエスト**
    ```json
    {
        "store_id": 1
    }
    ```
- **レスポンス**
    - 成功時: 予約成功メッセージ
    - 失敗時: エラーメッセージ

### 受け渡し
店舗側が予約確認によって表示された予約チケットに対し、受け渡しが完了した際に実行
- **URL**: `/store/reservation/delete`
- **メソッド**: `DELETE`
- **リクエスト**
    ```json
    {
        "reservation_id": 1
    }
    ```
- **レスポンス**
    - 成功時: 受け渡し成功メッセージ
    - 失敗時: エラーメッセージ

### ユーザー側の予約確認
ユーザーが自身の予約を確認するためのAPI
（user_idのみを本来想定しているが、reservation_idも追加しておく）
- **URL**: `/user/:id/reservation`
- **メソッド**: `GET`
- **リクエスト**
    ```json
    {
        "user_id": 1,
        "reservation_id": 1
    }
    ```
- **レスポンス**
    - 成功時: ユーザーの予約情報
    - 失敗時: エラーメッセージ

## 注意事項

- すべてのリクエストとレスポンスはJSON形式で行います。
- `user_id`などのユーザー情報は、セキュリティ上の理由から本番環境ではJWTトークンなどの認証機構を利用することを推奨します。
