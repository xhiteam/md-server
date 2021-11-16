package global

import (
	"time"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
}
