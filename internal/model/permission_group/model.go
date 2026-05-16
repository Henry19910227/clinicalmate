package permissiongroup

import "gorm.io/gorm"

type PermissionGroup struct {
	gorm.Model
	Name           string `gorm:"column:name;not null"`                        // 權限組名稱（如：系統管理員、普通用戶）
	PermissionName string `gorm:"column:permission_name;not null;default:''"` // 擁有的菜單名稱彙總（逗號分隔，供列表展示用）
}

// PermissionGroupPermission 是 permission_groups 與 permissions 的中間表
type PermissionGroupPermission struct {
	PermissionGroupID uint `gorm:"column:permission_group_id;not null;index"` // 權限組 ID
	PermissionID      uint `gorm:"column:permission_id;not null;index"`       // 菜單權限 ID
}
