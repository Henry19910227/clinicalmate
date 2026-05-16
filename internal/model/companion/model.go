package companion

import "gorm.io/gorm"

type Companion struct {
	gorm.Model
	Name   string `gorm:"column:name;not null"`   // 陪診師姓名
	Mobile string `gorm:"column:mobile;not null"`  // 手機號碼
	Avatar string `gorm:"column:avatar;not null"`  // 頭像 URL
	Sex    string `gorm:"column:sex;not null"`     // 性別（"1" 男 / "2" 女）
	Age    int    `gorm:"column:age;not null"`     // 年齡
	Active int    `gorm:"column:active;not null;default:1"` // 狀態（1: 啟用 / 0: 停用）
}
