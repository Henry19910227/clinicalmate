# domain-order

Notion 來源：https://www.notion.so/order-detail-35512ea11a3280348416cb0b29c2511a

## 對應資料表

`orders`

## 欄位說明

| 欄位 | Go 型別 | DB 型別 | 說明 |
|------|---------|---------|------|
| `out_trade_no` | `string` | `VARCHAR(64)` | 唯一訂單編號，對應 API 中的 `oid` 查詢參數 |
| `trade_state` | `string` | `VARCHAR(20)` | 訂單狀態：待支付 / 待服務 / 已完成 / 已取消；預設 `待支付` |
| `order_start_time` | `int64` | `BIGINT` | 下單時間戳（毫秒），對應 API 回傳的 `order_start_time` |
| `price` | `string` | `VARCHAR(20)` | 應付金額字串（如 `"100.00"`） |
| `service_name` | `string` | `VARCHAR(100)` | 預約服務名稱（如「就醫陪診」） |
| `start_time` | `time.Time` | `DATETIME` | 期望就診時間，對應 API 的 `starttime` |
| `tel` | `string` | `VARCHAR(20)` | 就診人備用聯絡電話 |
| `receive_address` | `string` | `VARCHAR(500)` | 接送地址，對應 API 的 `receiveAddress` |
| `demand` | `string` | `TEXT` | 備註需求，選填 |
| `code_url` | `string` | `VARCHAR(500)` | 支付二維碼連結，僅待支付時有值 |
| `client_name` | `string` | `VARCHAR(100)` | 就診人姓名，來自 API `client.name` |
| `client_mobile` | `string` | `VARCHAR(20)` | 就診人手機號碼，來自 API `client.mobile` |
| `user_id` | `uint` | `BIGINT UNSIGNED` | 下單用戶 ID，FK -> `users` |
| `hospital_id` | `uint` | `BIGINT UNSIGNED` | 就診醫院 ID，FK -> `hospitals` |
| `companion_id` | `uint` | `BIGINT UNSIGNED` | 指派陪診師 ID，FK -> `companions`；派單前為 0 |

## 設計決策

- `client_name` / `client_mobile` 扁平化儲存，不另建 client 子表，配合 API response 結構直接映射
- `companion_id` 允許 NULL，訂單建立時尚未派單，由後台派單後填入
- `price` 使用 `VARCHAR` 而非 `DECIMAL`，保持與 API 回傳格式一致，避免浮點精度問題
- `order_start_time` 使用 `int64` 毫秒時間戳，對應前端 JS timestamp 格式
- `hospital_name` 不存 model，透過 `hospital_id` join `hospitals` 取得
- companion 相關資訊（name / avatar / mobile）不存 model，透過 `companion_id` join `companions` 取得

## 索引

- `UNIQUE KEY uni_orders_out_trade_no` — 訂單編號全局唯一
- `INDEX idx_orders_user_id` — 依用戶查詢訂單列表
- `INDEX idx_orders_hospital_id` — 依醫院查詢訂單
- `INDEX idx_orders_companion_id` — 依陪診師查詢負責訂單
- `INDEX idx_orders_deleted_at` — GORM 軟刪除標準索引

## 訂單狀態流轉

```
待支付 -> 已取消（逾期未付）
待支付 -> 待服務（付款成功）
待服務 -> 已完成（服務結束）
待服務 -> 已取消（人工取消）
```
