# 各種APIの細かい想定をここに記述しています。

## ログイン処理（仮）
POST /login
front -> back ログインのユーザー情報 json
back -> front ログイン可否による情報

front -> back json
```
{
    "name": "user-name",
    "password": "user-password"
}
```
back -> front
ログイン成功 - 正常なリクエストを受け取ったことを示すクエリを返す
ログイン失敗 - 正常にログインできなかったことを示すクエリを返す

## ショップの一覧取得
GET /store
front -> back リクエストのみ
back -> front 取得した販売店情報の一覧を返す

front -> back
リクエストのみ
back -> front
```
{
    "store": "store_1",
    "store": "store_2",
        ... //省略
    "store": "store_n"
}
```

## ショップの弁当一覧
GET /store/{id}
front -> back 選択した販売店の名前
back -> front 取得した販売店の販売弁当の情報の一覧を返す
front -> back
```
{
    "store": "store_2"
}
```
back -> front
```
{
    "store": "store_2" {
        "menue": "bentou_1",
        "menue": "bentou_2",
        ...
        "menue": "bentou_n"
    }
}
```

## ショップの弁当選択
GET /menue/{id}
front -> back 選択したメニューの情報
back -> front 取得したメニューの購入可能な情報の返答
front -> back
```
{
    "store": "store_2",
    "menue": "bentou_1"
}
```
back -> front
```
{
    "store": "store_2",
    "menue": "bentou_1" {
        {
            "time": "12:00",
            "stock": "5"
        },
        {
            "time": "12:10",
            "stock": "3"
        },
        ...
        {
            "time": "hh:mm",
            "stock": "n"
        }
    }
}
```

## 購入可否
POST /payment
front -> back 購入情報
back -> front 購入可否の情報
front -> back
```
{
    "store": "store_2",
    "menue": "bentou_1",
    "time": "12:10",
    "count": "2"
}
```
back -> front
購入可能 - 購入可能なことを示すクエリを返す
購入不可能 - 購入不可能なことを示すクエリを返す

## 注文キャンセル
DELETE /cancel/{id}
// 未定

ショップ側の操作は一度後回し
