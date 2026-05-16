package hospital

import "gorm.io/gorm"

type Hospital struct {
	gorm.Model
	Name    string `gorm:"column:name;not null"`    // 醫院名稱
	Address string `gorm:"column:address;not null"` // 醫院地址
	Image   string `gorm:"column:image;not null"`   // 醫院圖片 URL
	Phone   string `gorm:"column:phone;not null"`   // 聯絡電話
}
