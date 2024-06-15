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

以下のテーブルを含むMariaDBデータベースを使用します。

### Userテーブル
- `id`: プライマリキー
- `name`: ユーザー名
- `mail`: ユーザーメールアドレス
- `created_at`: 作成日時
- `updated_at`: 更新日時

### Storeテーブル
- `id`: プライマリキー
- `name`: 店舗名
- `created_at`: 作成日時
- `updated_at`: 更新日時

### Menueテーブル
- `id`: プライマリキー
- `store_id`: Storeテーブルへの外部キー
- `name`: 弁当名
- `price`: 価格
- `description`: 説明
- `is_sold_out`: 売り切れフラグ
- `created_at`: 作成日時
- `updated_at`: 更新日時

### UserReservationテーブル
- `id`: プライマリキー
- `user_id`: Userテーブルへの外部キー
- `store_id`: Storeテーブルへの外部キー
- `menue_id`: Menueテーブルへの外部キー
- `reserv_time`: 予約時間
- `reserv_cnt`: 予約個数
- `is_recipt`: 受取済みフラグ
- `created_at`: 作成日時
- `updated_at`: 更新日時

### StoreReservationPolicyテーブル
- `id`: プライマリキー
- `store_id`: Storeテーブルへの外部キー
- `date`: 特定の日付
- `day_of_week`: 曜日
- `time_slot_interval`: 時間スロットの間隔（分）
- `max_reservations_per_slot`: 時間スロットあたりの最大予約数
- `created_at`: 作成日時
- `updated_at`: 更新日時

### StoreTimeSlotReservationテーブル
- `id`: プライマリキー
- `store_id`: Storeテーブルへの外部キー
- `date`: 予約日
- `time_slot`: 時間スロット
- `current_reservations`: 現在の予約数
- `created_at`: 作成日時
- `updated_at`: 更新日時

## エンドポイント

詳細なAPI仕様については、`API-README.md`を参照してください。

