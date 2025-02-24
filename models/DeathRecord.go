package models

import (
	"time"

	"gorm.io/gorm"
)

type DeathRecord struct {
	UUID       string `gorm:"primaryKey;unique"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Date       time.Time
	CattleUUID string `gorm:"index"`
	Cause      string
	Cattle     *Cattle `gorm:"foreignKey:CattleUUID;references:UUID"`
}
