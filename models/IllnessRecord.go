package models

import (
	"time"

	"gorm.io/gorm"
)

type IllnessRecord struct {
	UUID                string `gorm:"primaryKey;unique"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	StartDate           time.Time
	EndDate             time.Time
	CattleUUID          string `gorm:"index"`
	Name                string
	Cattle              *Cattle `gorm:"foreignKey:CattleUUID;references:UUID"`
	AreAntibioticsUsing bool
	BlocksMilking       bool
}
