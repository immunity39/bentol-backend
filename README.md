# 弁当予約販売システム

このプロジェクトは、GoのGinフレームワークとMariaDBを使用して構築された弁当予約販売システムです。ユーザーは販売店から弁当を予約し、指定した時間に受け取ることができます。
予約は一定時間の購入個数上限に引っかからないときに予約確定されます。

## 動作仕様
### 店舗
1. 弁当の販売者は店舗の名前、パスワード、メールを登録することで新規店舗の登録ができます。
2. 店舗登録が完了している販売者はloginページからログインができます。
3. 店舗の登録が完了後、店舗の基本販売スケジュールポリシーを設定します。
スケジュールポリシーは、店舗の開店閉店時間や一定間隔で区切る時間の幅の設定、定間隔時間の予約可能な最大個数などを設定します。設定は曜日ごとに設定が可能です。
4. 店舗のスケジュールポリシーが完了すると、cron jobによって毎日午前三時に一週間先までの店舗の予約スケジュールDBが作成されます。
5. 特定日においては、店舗側のスケジュールポリシーを別のものにすることができます。これは、休みの日やイベント日のような日に対応するためです。
特定日の予約スケジュール情報をもとに毎日午前四時に一週間先までの特定予約情報が店舗の予約スケジュールDBに作成（更新）されます。
6. また、店舗の経営者はメニューを追加、更新することが可能です。

### ユーザー
1. ユーザーはユーザー名、パスワード、メールを登録することで新規ユーザー登録ができます。
2. ユーザー登録が完了しているユーザーはloginページからログインができます。
3. ログイン完了後、ユーザーは店舗一覧（ホームページ）から店舗を選択することができます。
4. 店舗を選択すると、店舗の販売メニューを見ること、そして販売メニューの選択を行うことができます。
5. 店舗の販売メニューを選択すると、予約時間と予約個数を指定することで予約を行うことができます。店舗の予約スケジュール情報をもとに予約可能かどうかの判断がなされたのち、可能な場合は支払いのページに飛びます。
6. 支払いはpaypay apiを利用した前払いとなります。
7. 受け取りが完了すると、再度注文を行うことができるようになります。

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

### StoreBasicReservationPolicy

- `id`: ポリシーID
- `store_id`: 店舗ID
- `day_of_week`: 曜日
- `time_slot_interval`: 時間スロット間隔
- `max_reservations_per_slot`: スロットごとの最大予約数
- `store_start_time`: 店舗の開店時間
- `store_end_time`: 店舗の閉店時間
- `created_at`: 作成日時
- `updated_at`: 更新日時

### StoreSpecificReservationPolicy

- `id`: ポリシーID
- `store_id`: 店舗ID
- `date`: 日付
- `time_slot_interval`: 時間スロット間隔
- `max_reservations_per_slot`: スロットごとの最大予約数
- `store_start_time`: 店舗の開店時間
- `store_end_time`: 店舗の閉店時間
- `created_at`: 作成日時
- `updated_at`: 更新日時

### StoreReservationSchedule

- `id`: スケジュールID
- `store_id`: 店舗ID
- `date`: 日付
- `time_slot_interval`: 時間スロット間隔
- `max_reservations_per_slot`: スロットごとの最大予約数
- `store_start_time`: 店舗の開店時間
- `store_end_time`: 店舗の閉店時間
- `created_at`: 作成日時
- `updated_at`: 更新日時

### StoreVendor

- `id`: 管理者ID
- `store_id`: 店舗ID
- `name`: 販売者名
- `email`: メールアドレス
- `password`: パスワード
- `created_at`: 作成日時
- `updated_at`: 更新日時

## API仕様

詳細は `API-README.md` を参照してください。

