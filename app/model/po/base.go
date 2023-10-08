package po

import "time"

type BaseTimeColumns struct {
	CreateAt time.Time `gorm:"column:create_at;autoCreateTime"`
	UpdateAt time.Time `gorm:"column:update_at;autoUpdateTime"`
}

type BaseId struct {
	Id uint `gorm:"column:id"`
}
