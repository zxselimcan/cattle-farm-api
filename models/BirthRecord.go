package models

import (
	"time"

	"gorm.io/gorm"
)

type BirthRecord struct {
	UUID          string `gorm:"primaryKey;unique"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Date          time.Time
	MotherUUID    string  `gorm:"index"`
	ChildUUID     string  `gorm:"primaryKey"`
	Child         *Cattle `gorm:"foreignKey:ChildUUID;references:UUID"`
	Mother        *Cattle `gorm:"foreignKey:MotherUUID;references:UUID"`
	ChildrenCount uint
	Type          string
}
