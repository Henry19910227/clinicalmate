package order

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OutTradeNo     string          `gorm:"column:out_trade_no;not null;uniqueIndex"` // 唯一訂單編號
	TradeState     string          `gorm:"column:trade_state;not null;default:'待支付'"` // 訂單狀態（待支付 / 待服務 / 已完成 / 已取消）
	OrderStartTime int64           `gorm:"column:order_start_time;not null"`         // 下單時間戳（毫秒）
	Price          string          `gorm:"column:price;not null"`                    // 應付金額（如 "100.00"）
	ServiceName    string          `gorm:"column:service_name;not null"`             // 預約服務名稱
	StartTime      time.Time       `gorm:"column:start_time;not null"`               // 期望就診時間
	Tel            string          `gorm:"column:tel;not null"`                      // 就診人備用聯絡電話
	ReceiveAddress string          `gorm:"column:receive_address;not null"`          // 接送地址
	Demand         string          `gorm:"column:demand"`                            // 備註需求（選填）
	CodeURL        string          `gorm:"column:code_url"`                          // 支付二維碼連結（待支付時使用）
	ClientName     string          `gorm:"column:client_name;not null"`              // 就診人姓名
	ClientMobile   string          `gorm:"column:client_mobile;not null"`            // 就診人手機號碼
	UserID         uint            `gorm:"column:user_id;not null"`                  // 下單用戶 ID（FK -> users）
	HospitalID     uint            `gorm:"column:hospital_id;not null"`              // 就診醫院 ID（FK -> hospitals）
	CompanionID    uint            `gorm:"column:companion_id"`                      // 指派陪診師 ID（FK -> companions，派單後填入）
}
