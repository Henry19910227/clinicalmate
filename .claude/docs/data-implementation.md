## 資料層架構

**Request flow（資料方向）**: `Store → Repository → DB`

```
internal/
  model/              # GORM model structs（對應資料庫表格）
  repository/<domain>/
    interface.go      # Repository interface 定義
    repository.go     # GORM 單表查詢實作
  store/<domain>/
    interface.go      # Store interface 定義
    store.go          # 跨 repo transaction + domain↔DB model mapping
scripts/
  sql/                # 建表 SQL 腳本（如果目錄不存在，執行任務前先建立）
```

## MCP 工具

### Notion

- `mcp__notionApi__API-post-search`：搜尋相關頁面或資料庫
- `mcp__notionApi__API-get-page`：讀取特定頁面內容
- `mcp__notionApi__API-get-database`：讀取資料庫結構
- `mcp__notionApi__API-post-database-query`：查詢資料庫記錄
- `mcp__notionApi__API-get-block-children`：讀取頁面區塊內容

從文件中擷取：欄位名稱、資料型別、約束條件（NOT NULL、UNIQUE、INDEX）、外鍵關係、業務說明。

### MySQL

- `mcp__mysql__list_tables`：列出所有資料表
- `mcp__mysql__describe_table`：查看表格欄位定義
- `mcp__mysql__query`：執行任意 SQL 查詢（用於 SHOW CREATE TABLE、查詢索引等）

比對 Notion 文件與實際資料庫，找出差異。

## GORM Model Struct 規範

檔案位置：`internal/model/<domain>.go`

**命名規範**：
- Struct 名稱：PascalCase，對應資料表的單數名（如 `users` → `User`）
- 欄位名稱：PascalCase（如 `CreatedAt`）
- GORM tag：`gorm:"column:snake_case_name"`
- JSON tag：`json:"snake_case_name"`

**範本**：
```go
package model

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
    Name      string         `gorm:"column:name;not null"     json:"name"`
    Email     string         `gorm:"column:email;uniqueIndex;not null" json:"email"`
    CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"  json:"created_at"`
    UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"  json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"           json:"deleted_at"`
}

func (User) TableName() string {
    return "users"
}
```

**與 Notion 文件同步原則**：
- 文件新增欄位 → struct 新增對應欄位 + 更新 SQL
- 文件修改型別或約束 → struct 同步修改 + 更新 SQL
- 不可自行新增文件未定義的欄位

## 建表 SQL 腳本規範

檔案位置：`scripts/sql/<domain>_<table>.sql`（如 `scripts/sql/patient_users.sql`）

**規範**：
- 使用 `CREATE TABLE IF NOT EXISTS`
- 欄位順序：PK → 業務欄位 → `created_at` → `updated_at` → `deleted_at`（軟刪除）
- 字符集：`CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci`
- 引擎：`ENGINE=InnoDB`
- 索引單獨 `ALTER TABLE` 或寫在建表語句中
- 加上 `-- Notion doc: <page_url>` 的注釋，追蹤文件來源

**範本**：
```sql
-- Notion doc: https://notion.so/...
CREATE TABLE IF NOT EXISTS `users` (
    `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(100)    NOT NULL,
    `email`      VARCHAR(255)    NOT NULL,
    `created_at` DATETIME(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` DATETIME(3)     NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uq_users_email` (`email`),
    KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

## Repository 層實作規範

**interface.go 範本**：
```go
package <domain>

import "clinicalmate/internal/model"

type Repository interface {
    Create(entity *model.User) error
    FindByID(id uint) (*model.User, error)
    FindAll() ([]*model.User, error)
    Update(entity *model.User) error
    Delete(id uint) error
}
```

**repository.go 範本**：
```go
package <domain>

import (
    "clinicalmate/internal/model"
    "gorm.io/gorm"
)

type repository struct {
    db *gorm.DB
}

func New(db *gorm.DB) Repository {
    return &repository{db: db}
}

func (r *repository) Create(entity *model.User) error {
    return r.db.Create(entity).Error
}

func (r *repository) FindByID(id uint) (*model.User, error) {
    var entity model.User
    if err := r.db.First(&entity, id).Error; err != nil {
        return nil, err
    }
    return &entity, nil
}
```

**限制**：
- 每個方法只操作單一資料表
- 不寫業務邏輯（條件判斷屬於 service）
- 不管理 transaction（由 store 負責）

## Store 層實作規範

**interface.go 範本**：
```go
package <domain>

type Store interface {
    CreateUser(name, email string) error
    GetUserByID(id uint) (*UserDomain, error)
}

type UserDomain struct {
    ID    uint
    Name  string
    Email string
}
```

**store.go 範本**：
```go
package <domain>

import (
    "clinicalmate/internal/model"
    domainRepo "clinicalmate/internal/repository/<domain>"
    "gorm.io/gorm"
)

type store struct {
    db         *gorm.DB
    domainRepo domainRepo.Repository
}

func New(db *gorm.DB, domainRepo domainRepo.Repository) Store {
    return &store{db: db, domainRepo: domainRepo}
}

func (s *store) CreateUser(name, email string) error {
    return s.db.Transaction(func(tx *gorm.DB) error {
        entity := &model.User{Name: name, Email: email}
        return s.domainRepo.Create(entity)
    })
}

func (s *store) GetUserByID(id uint) (*UserDomain, error) {
    entity, err := s.domainRepo.FindByID(id)
    if err != nil {
        return nil, err
    }
    return &UserDomain{ID: entity.ID, Name: entity.Name, Email: entity.Email}, nil
}
```

**限制**：
- Domain model（Store 的 interface 型別）不暴露 GORM struct
- Transaction 邏輯在 Store，不下放到 Repository
- 不含業務規則（如「用戶只能有一個帳號」屬於 Service）
