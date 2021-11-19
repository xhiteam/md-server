package global

type GVA_MODEL struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"` // 主键ID
	CreatedAt LocalTime // 创建时间
	UpdatedAt LocalTime // 更新时间
}
