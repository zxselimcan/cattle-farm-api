package models

import (
	"time"

	"gorm.io/gorm"
)

type MilkingRecord struct {
	UUID       string `gorm:"primaryKey;unique"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Date       time.Time
	MilkAmount uint
	CattleUUID string  `gorm:"index"`
	Cattle     *Cattle `gorm:"foreignKey:CattleUUID;references:UUID"`
}
