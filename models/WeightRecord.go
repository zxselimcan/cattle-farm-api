package models

import (
	"time"

	"gorm.io/gorm"
)

type WeightRecord struct {
	UUID       string `gorm:"primaryKey;unique"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Date       time.Time
	CattleUUID string `gorm:"index"`
	Weight     uint
	Cattle     *Cattle `gorm:"foreignKey:CattleUUID;references:UUID"`
}
