package permission

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name      string `gorm:"column:name;not null;uniqueIndex"`     // 路由名稱（唯一，對應前端 router name）
	Label     string `gorm:"column:label;not null"`               // 顯示名稱（如「菜单管理」）
	Path      string `gorm:"column:path;not null;default:''"`     // 路由路徑
	Component string `gorm:"column:component;not null;default:''"` // 前端組件映射
	Icon      string `gorm:"column:icon;not null;default:''"`     // 圖標
	ParentID  uint   `gorm:"column:parent_id;not null;default:0"` // 父節點 ID（0 代表頂層）
	Disabled  int    `gorm:"column:disabled;not null;default:0"`  // 是否停用（0: 啟用 / 1: 停用）
}
