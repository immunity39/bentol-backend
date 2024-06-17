# 弁当予約販売システム

このプロジェクトは、GoのGinフレームワークとMariaDBを使用して構築された弁当予約販売システムです。ユーザーは販売店から弁当を予約し、指定した時間に受け取ることができます。

## プロジェクト構成

- `main.go`: アプリケーションのエントリーポイント
- `config/db.go`: データベース接続およびテーブルマイグレーション
- `models/`: データベースモデル
- `controllers/`: APIコントローラー
- `routes/`: ルート設定
- `utils/`: ユーティリティ関数（将来的に追加予定）

## データベース構造

### User

- `id`: ユーザーID
- `name`: ユーザー名
- `mail`: メールアドレス
- `created_at`: 作成日時
- `updated_at`: 更新日時

### Store

- `id`: 店舗ID
- `name`: 店舗名
- `created_at`: 作成日時
- `updated_at`: 更新日時

### Menue

- `id`: メニューID
- `store_id`: 店舗ID
- `name`: メニュー名
- `price`: 価格
- `description`: 説明
- `is_sold_out`: 売り切れフラグ
- `created_at`: 作成日時
- `updated_at`: 更新日時

### UserReservation

- `id`: 予約ID
- `user_id`: ユーザーID
- `store_id`: 店舗ID
- `menue_id`: メニューID
- `reserv_time`: 予約受け取り時間
- `reserv_cnt`: 予約個数
- `is_recipt`: 受け取り済みフラグ
- `created_at`: 作成日時
- `updated_at`: 更新日時

### StoreReservationPolicy

- `id`: ポリシーID
- `store_id`: 店舗ID
- `date`: 日付
- `day_of_week`: 曜日
- `time_slot_interval`: 時間スロット間隔
- `max_reservations_per_slot`: スロットごとの最大予約数
- `created_at`: 作成日時
- `updated_at`: 更新日時

### StoreTimeSlotReservation

- `id`: タイムスロット予約ID
- `store_id`: 店舗ID
- `date`: 日付
- `time_slot`: 時間スロット
- `current_reservations`: 現在の予約数
- `created_at`: 作成日時
- `updated_at`: 更新日時

### StoreAdmin

- `id`: 管理者ID
- `store_id`: 店舗ID
- `email`: メールアドレス
- `password`: パスワード
- `created_at`: 作成日時
- `updated_at`: 更新日時

## API仕様

詳細は `API-README.md` を参照してください。
