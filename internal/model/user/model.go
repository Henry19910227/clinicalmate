package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"column:user_name;not null"` // 暱稱或顯示名稱
	Avatar   string `gorm:"column:avatar;not null"`    // 大頭照 URL
	Mobile   string `gorm:"column:mobile;not null;uniqueIndex"` // 用戶綁定的手機號（全系統唯一）
	Password string `gorm:"column:password;not null"`          // 登入密碼（雜湊儲存）
}
