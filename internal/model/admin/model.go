package admin

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name          string `gorm:"column:name;not null"`              // 管理員名稱
	Mobile        string `gorm:"column:mobile;not null;uniqueIndex"` // 手機號碼（登入識別，全系統唯一）
	PermissionsID int    `gorm:"column:permissions_id;not null"`    // 所屬權限組 ID
	Active        int    `gorm:"column:active;not null;default:1"`  // 狀態（1: 正常 / 0: 失效）
	Password      string `gorm:"column:password;not null"`          // 登入密碼（雜湊儲存）
}
